package web

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/cgrice/dnslookup/dns"
)

type ResultsPageData struct {
    Query   string
    Results []dns.Record
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./web/templates/layout.html", "./web/templates/home.html"))

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
	
	fmt.Println(nameserver, recordtype)


	results := dns.Query(vars["domain"], recordtype, nameserver)

	t := template.Must(template.ParseFiles("./web/templates/layout.html", "./web/templates/home.html"))

	fmt.Println(results)

	t.Execute(w, ResultsPageData{
		Results: results,
		Query: vars["domain"],
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
		redirectURL = redirectURL + "/" + r.FormValue("recordtype")
	}

	// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

	http.Redirect(w, r, redirectURL, 301)
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

    r.HandleFunc("/", home)
	r.HandleFunc("/query", query).Methods("POST")
	r.HandleFunc("/{domain}@{nameserver}/{type}", lookup)
	r.HandleFunc("/{domain}/{type}", lookup)
	r.HandleFunc("/{domain}@{nameserver}", lookup)
	r.HandleFunc("/{domain}", lookup)
	
	return r
}