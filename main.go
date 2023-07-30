package main

import (
	"fmt"
	"net/http"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to the world of Klckr</h1>")
}

func main() {
	http.HandleFunc("/", BasicHandler)
	fmt.Println("Server will listen on port : 4949")
	http.ListenAndServe(":4949", nil)
}
