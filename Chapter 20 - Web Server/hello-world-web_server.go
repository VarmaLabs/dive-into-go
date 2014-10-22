package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var chttp = http.NewServeMux()

func main() {

	chttp.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/", HomeHandler) // homepage

	log.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "HomeHandler")
	}
}

/*

Output
------
It will start a web service in port 8080 and will respond to a GET request to localhost:8080

Notes
-----

"net/http" is a package which provides HTTP client and server implementations.
http.ListenAndServe accepts an address (with port) and a handler will starts an HTTP server which listen to the port passed.
The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

*/
