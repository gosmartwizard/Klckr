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

func parseExecuteTemplate(w http.ResponseWriter, tplPath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	parseExecuteTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	parseExecuteTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	parseExecuteTemplate(w, tplPath)
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
