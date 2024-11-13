package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":       "available",
		"environmennt": app.config.env,
		"version":      version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// js := `{"status":"available", "environmet":%q,"version":%q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	//make it easy to read in the terminal
}
