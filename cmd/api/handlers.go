// Filename: cmd/api/handlers.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abelwhite/orange/internal/data"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a school...")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "School displayed...")
	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
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
	err = app.WriteJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

}
