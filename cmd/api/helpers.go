// Filename: cmd/api/helpers.go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{} //envelope is a type that will contain school infos //container inside a container

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid Id Parameter")
	}
	return id, nil
}

func (app *application) WriteJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//convert the data to json format
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n') //if no error
	//add any headers that were sent
	for key, value := range headers {
		w.Header()[key] = value
	}

	//add header info
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(js))
	return nil

}

// to help signify the error from the api
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, destination interface{}) error {
	//specify the max size of our JSON request body
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	//we will try to decode a JSON request
	dec := json.NewDecoder(r.Body) //created the decoder
	dec.DisallowUnknownFields()
	//start the decoding
	err := dec.Decode(destination)
	if err != nil {
		//something went wrong
		//lets find out the type of error
		var syntaxError *json.SyntaxError
		var unMarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		//check for max bytes
		var maxBytesError *http.MaxBytesError

		//lets check for the type of decode error
		switch {
		case errors.As(err, &syntaxError): //we use As if we looking for the type of error
			return fmt.Errorf("body contains badly-formed JSON(at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF): //we use Is if we looking for a specific type of error
			return fmt.Errorf("body contains badly-formed JSON")

		case errors.As(err, &unMarshalTypeError):
			if unMarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q)", unMarshalTypeError.Field)
			}
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", unMarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
			//unmappable field/ not-existent field
		case strings.HasPrefix(err.Error(), "json: unkown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json : unkown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	//Lets call the decorder to check if they are any trailing json objects
	err = dec.Decode(&struct{}{}) //decode functions takes the json request and dumps it somewhere else
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil

}
