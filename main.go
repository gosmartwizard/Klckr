package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/gosmartwizard/Klckr/controllers"
	"github.com/gosmartwizard/Klckr/views"
)

func parseExecuteTemplate(w http.ResponseWriter, tplPath string) {

	tpl, err := views.Parse(tplPath)
	if err != nil {
		log.Printf("Parse template failed with error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
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
	/*router := chi.NewRouter()

	//router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFound(notFoundHandler)
	//router.Get("/users/{userID}", getParamsHandler)
	router.With(middleware.Logger).Get("/users/{userID}", getParamsHandler)

	fmt.Println("Server will listen on port : 4949")

	http.ListenAndServe(":4949", router) */

	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :4949...")
	http.ListenAndServe(":4949", r)
}
