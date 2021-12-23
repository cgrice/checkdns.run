package web

import (
	"fmt"
	"embed"
	"time"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/cgrice/dnslookup/dns"
)

type ResultsPageData struct {
    Query   string
	Nameserver string
    Results []dns.Record
	Latency time.Duration
}

//go:embed templates/*.html
var templateData embed.FS

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templateData, "templates/layout.html", "templates/home.html"))
    t.Execute(w, ResultsPageData{})
}

func lookup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nameserver := "8.8.8.8"
	recordtype := "A"

	if(vars["nameserver"] != "") {
		nameserver = vars["nameserver"]		
	}

	if(vars["type"] != "") {
		recordtype = vars["type"]		
	}
	
	results, latency := dns.Query(vars["domain"], recordtype, nameserver)

	t := template.Must(template.ParseFS(templateData, "templates/layout.html", "templates/home.html"))

	t.Execute(w, ResultsPageData{
		Results: results,
		Query: vars["domain"],
		Nameserver: nameserver,
		Latency: latency.Round(time.Millisecond),
	})
}

func query(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	redirectURL := "/" + r.FormValue("domain")

	if(r.FormValue("nameserver") != "") {
		redirectURL = redirectURL + "@" + r.FormValue("nameserver")
	}

	if(r.FormValue("recordtype") != "") {
		redirectURL = r.FormValue("recordtype") + "/" + redirectURL
	}

	http.Redirect(w, r, redirectURL, 301)
}

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