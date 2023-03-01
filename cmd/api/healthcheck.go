// Filename: .cmd/api/healthchech.go
package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//js := `{"status": "available", "environment" :%q, "version": %q}`
	//js = fmt.Sprintf(js, app.config.env, version)

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.WriteJSON(w, http.StatusOK, data, nil) //sen ok to say everything is okay. We send w caz it semds stuff to the browser.
	if err != nil {
		return
	}

}
