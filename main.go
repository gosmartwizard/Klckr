package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the world of Klckr</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:nrkr@klckr.io\">nrkr@klckr.io</a>.</p>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Requested path : " + r.URL.Path + " is not valid and Page not found"
	http.Error(w, msg, http.StatusNotFound)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/contact":
		contactHandler(w, r)
	case "/":
		homeHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Server will listen on port : 4949")
	http.ListenAndServe(":4949", nil)
}
