package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

func NewHealthzHandler() http.Handler {
	route := chi.NewRouter()
	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			slog.Error(err.Error())
		}

		slog.Info("Server connection fine")
	})

	return route
}
