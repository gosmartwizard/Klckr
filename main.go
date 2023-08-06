package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	t, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Failed to parse template file : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Failed to execute template : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

func getParamsHandler(w http.ResponseWriter, r *http.Request) {
	v := chi.URLParam(r, "userID")
	msg := fmt.Sprintf("<h1>UserID : %v </h1>", v)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, msg)
}

func main() {
	router := chi.NewRouter()

	//router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFound(notFoundHandler)
	//router.Get("/users/{userID}", getParamsHandler)
	router.With(middleware.Logger).Get("/users/{userID}", getParamsHandler)

	fmt.Println("Server will listen on port : 4949")

	http.ListenAndServe(":4949", router)
}
