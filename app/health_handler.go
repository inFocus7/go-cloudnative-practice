package app

import "net/http"

func HandleLive(w http.ResponseWriter, _ *http.Request) {
	writeHealth(w, true)
}

func (app *App) HandleReady(w http.ResponseWriter, r *http.Request) {
	if err := app.db.DB().Ping(); err != nil {
		app.Logger().Fatal().Err(err).Msg("")
		writeHealth(w, false)
		return
	}

	writeHealth(w, true)
}

func writeHealth(w http.ResponseWriter, healthy bool) {
	w.Header().Set("Content-Type", "text/plain")
	if healthy {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte("."))
}
