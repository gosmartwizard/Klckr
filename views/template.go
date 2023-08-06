package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Parse(path string) (Template, error) {

	htmlTpl, err := template.ParseFiles(path)
	if err != nil {
		return Template{}, fmt.Errorf("error parsing the template files : %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Failed to execute template : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
