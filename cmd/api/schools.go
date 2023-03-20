// Filename: cmd/api/handlers.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abelwhite/orange/internal/data"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//create a struct to hold a school that will be provided to us
	//via the request
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website,omitempty"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}
	//Initialize a new json.Decoder instance
	err := json.NewDecoder(r.Body).Decode(&input) //we take r.Body and decode it into input
	//take it from json and decode it so it can be sent to db.
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	//Print the request
	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "School displayed...")
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//fmt.Fprintf(w, "Show details of school %d \n ", id)
	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "University of Belmopan",
		Level:     "University",
		Contact:   "Abel Blanco",
		Phone:     "323-4545",
		Website:   "https://uob.edu.bz",
		Address:   "17 Apple Avenue",
		Mode:      []string{"blended", "online", "face-to-face"},
		Version:   1,
	}
	err = app.WriteJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err) //
		return
	}

}
