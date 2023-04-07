package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func (w *WebServer) AddHandler(route string, handle http.HandlerFunc) {
	w.Handlers[route] = handle
}

func (w *WebServer) Start() {
	w.Router.Use(middleware.Logger))
	for path, handler := range w.Handlers {
		w.Router.HandleFunc(path, handler)
	}
	http.ListenAndServe(w.WebServerPort, w.Router)
}
