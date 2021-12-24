package web

import (
	"embed"
	"fmt"
	"sync"
	"github.com/cgrice/dnslookup/dns"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type ResultsPageData struct {
	Query      string
	Nameserver string
	Servers    []dns.Server
	RecordType string
	ShowAdvanced bool
	ShowDetail bool
	Results    []dns.Record
	Latency    time.Duration
	Permalink string
}

type PropogationResult struct {
	Result dns.Record
	Server dns.Server
	Found bool
	Latency time.Duration
}

type PropogationPageData struct {
	Query      string
	Results    []PropogationResult
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
		Servers: dns.GetServers(),
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

	t := template.Must(template.ParseFS(templateData, "templates/layout.html", "templates/home.html"))

	pageData := getPageData(r)

	pageData.Results = results
	pageData.RecordType = recordtype
	pageData.Query = vars["domain"]
	pageData.Nameserver = nameserver
	pageData.Latency = latency.Round(time.Millisecond)

	t.Execute(w, pageData)
}

func propogation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var wg sync.WaitGroup

	propogationResults := make(chan PropogationResult, 100)

	for _, server := range dns.GetServers() {
		wg.Add(1)

		server := server

		go func() {
			defer wg.Done()
			results, latency := dns.Query(vars["domain"], "A", server.Address)

			if (len(results) > 0) {
				propogationResults <- PropogationResult{
					Result: results[0],
					Server: server,
					Found: true,
					Latency: latency,
				}
			} else {
				propogationResults <- PropogationResult{
					Server: server,
					Found: false,
					Latency: latency,
				}
			}
		}()
	}

	wg.Wait()
	close(propogationResults)

	t := template.Must(template.ParseFS(templateData, "templates/layout.html", "templates/propogation.html"))

	results := make([]PropogationResult, 0)
	for result := range propogationResults {
		results = append(results, result)
	}

	t.Execute(w, PropogationPageData{
		Query: vars["domain"],
		Results: results,
	})
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

func checkPropogation(w http.ResponseWriter, r *http.Request) {
	redirectURL := "/" + r.FormValue("domain") + "/propogation"
	http.Redirect(w, r, redirectURL, 301)
}