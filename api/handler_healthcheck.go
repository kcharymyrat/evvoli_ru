package api

import (
	"net/http"
)

func (app *ApiApplication) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":     "available",
		"environmet": app.Config.Env,
		"version":    version,
	}

	err := app.writeJson(w, 200, data, nil)
	if err != nil {
		app.Logger.Println(err)
		http.Error(w, "The server error", http.StatusInternalServerError)
	}
}
