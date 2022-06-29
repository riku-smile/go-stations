package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		handler.NewHealthzHandler().ServeHTTP(w, r)
	})

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		svc := service.NewTODOService(todoDB)
		handler.NewTODOHandler(svc).ServeHTTP(w, r)
	})

	return mux
}
