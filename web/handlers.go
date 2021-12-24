package web

import (
	"embed"
	"fmt"
	"github.com/cgrice/dnslookup/dns"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type ResultsPageData struct {
	Query      string
	Nameserver string
	RecordType string
	ShowAdvanced bool
	ShowDetail bool
	Results    []dns.Record
	Latency    time.Duration
	Permalink string
}

//go:embed templates/*.html
var templateData embed.FS

func getPageData(r *http.Request) ResultsPageData {
	showAdvanced, _ := r.Cookie("showAdvanced")
	showDetail, _ := r.Cookie("showDetail")

	permalink := r.URL.Path

	if (showAdvanced == nil) {
		showAdvanced = &http.Cookie{
			Name:   "showAdvanced",
			Value:  "false",
			MaxAge: 300,
		}
	}

	if (showDetail == nil) {
		showDetail = &http.Cookie{
			Name:   "showDetail",
			Value:  "false",
			MaxAge: 300,
		}
	}

	return ResultsPageData{
		ShowAdvanced: showAdvanced.Value == "true",
		ShowDetail: showDetail.Value == "true",
		Permalink: permalink,
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templateData, "templates/layout.html", "templates/home.html"))
	t.Execute(w, getPageData(r))
}

func lookup(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)

	nameserver := "8.8.8.8"
	recordtype := "A"

	if vars["nameserver"] != "" {
		nameserver = vars["nameserver"]
	}

	if vars["type"] != "" {
		recordtype = vars["type"]
	}

	results, latency := dns.Query(vars["domain"], recordtype, nameserver)

	fmt.Println(results)

	t := template.Must(template.ParseFS(templateData, "templates/layout.html", "templates/home.html"))

	pageData := getPageData(r)

	pageData.Results = results
	pageData.RecordType = recordtype
	pageData.Query = vars["domain"]
	pageData.Nameserver = nameserver
	pageData.Latency = latency.Round(time.Millisecond)

	t.Execute(w, pageData)
}

func query(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	redirectURL := "/" + r.FormValue("domain")

	if r.FormValue("nameserver") != "" {
		redirectURL = redirectURL + "@" + r.FormValue("nameserver")
	}

	if r.FormValue("recordtype") != "" {
		redirectURL = r.FormValue("recordtype") + "/" + redirectURL
	}

	http.Redirect(w, r, redirectURL, 301)
}
