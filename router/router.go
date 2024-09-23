package router

import (
	"github.com/gorilla/mux"
	"github.com/xiaoxuan6/sensitive-api/handlers"
	"net/http"
)

func Register() *mux.Router {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllow)

	r.HandleFunc("/", handlers.Index).Methods("GET")

	g := r.PathPrefix("/sensitive").Subrouter()
	g.HandleFunc("/filter", handlers.Filter).Methods("POST")
	g.HandleFunc("/findall", handlers.FindAll).Methods("POST")
	g.HandleFunc("/replace", handlers.Replace).Methods("POST")
	g.HandleFunc("/validate", handlers.Validate).Methods("POST")

	return r
}
