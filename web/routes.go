package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	r.HandleFunc("/favicon.ico", faviconHandler)
	r.HandleFunc("/", home)
	r.HandleFunc("/query", query).Methods("POST")

	r.HandleFunc("/check-propagation", checkPropagation).Methods("POST")
	r.HandleFunc("/{type}/{domain}@{nameserver}/propagation", propagation)
	r.HandleFunc("/{type}/{domain}/propagation", propagation)
	r.HandleFunc("/{domain}@{nameserver}/propagation", propagation)
	r.HandleFunc("/{domain}/propagation", propagation)

	r.HandleFunc("/{type}/{domain}@{nameserver}", lookup)
	r.HandleFunc("/{type}/{domain}", lookup)
	r.HandleFunc("/{domain}@{nameserver}", lookup)

	r.HandleFunc("/{domain}", lookup)

	return r
}
