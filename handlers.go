package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
