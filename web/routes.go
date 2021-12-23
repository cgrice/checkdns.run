package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	r.HandleFunc("/", home)
	r.HandleFunc("/query", query).Methods("POST")
	r.HandleFunc("/{type}/{domain}@{nameserver}", lookup)
	r.HandleFunc("/{type}/{domain}", lookup)
	r.HandleFunc("/{domain}@{nameserver}", lookup)
	r.HandleFunc("/{domain}", lookup)

	return r
}
