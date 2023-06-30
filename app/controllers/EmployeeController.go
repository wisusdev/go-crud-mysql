package controllers

import (
	"html/template"
	"main/app/models"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("resources/views/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	employee := models.GetAll()
	tmpl.ExecuteTemplate(w, "Index", employee)
}

func Show(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	emp := models.GetById(nId)
	tmpl.ExecuteTemplate(w, "Show", emp)

}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	emp := models.GetById(nId)
	tmpl.ExecuteTemplate(w, "Edit", emp)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")

		models.Create(name, city)
	}

	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")

		models.UpdateById(name, city, id)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	emp := r.URL.Query().Get("id")
	models.DeleteById(emp)
	http.Redirect(w, r, "/", 301)
}
