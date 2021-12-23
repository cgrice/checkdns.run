package main

import (
	"github.com/cgrice/dnslookup/web"
	"net/http"
)

func main() {
	router := web.GetRouter()

	http.ListenAndServe(":8080", router)
}
