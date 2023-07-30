package main

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the world of Klckr</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:nrkr@klckr.io\">nrkr@klckr.io</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
  <ul>
	<li>
	  <b>Is there a free version?</b>
	  Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
	  <b>What are your support hours?</b>
	  We have support staff answering emails 24/7, though response
	  times may be a bit slower on weekends.
	</li>
	<li>
	  <b>How do I contact support?</b>
	  Email us - <a href="mailto:support@klckr.com">support@klckr.com</a>
	</li>
  </ul>
  `)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Requested path : " + r.URL.Path + " is not valid and Page not found"
	http.Error(w, msg, http.StatusNotFound)
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/contact":
		contactHandler(w, r)
	case "/":
		homeHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Server will listen on port : 4949")
	http.ListenAndServe(":4949", router)
}
