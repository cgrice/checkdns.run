package main

import (
	"net/http"
	"github.com/cgrice/dnslookup/web"
)

func main() {
	router := web.GetRouter()

    http.ListenAndServe(":8080", router)
}

// import (
//     "fmt"
//     "log"
//     "net/http"
// )


// type Record struct {
// 	query string
// 	result string
// 	ttl int
// 	nameserver string
// 	recordtype string
// 	latency int
// }

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
//     http.HandleFunc("/", handler)
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }