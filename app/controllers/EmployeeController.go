package controllers

import (
	"html/template"
	"log"
	"main/core"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

var tmpl = template.Must(template.ParseGlob("resources/views/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := core.DbConnMySQL()
	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}

	defer db.Close()

	tmpl.ExecuteTemplate(w, "Index", res)

}

func Show(w http.ResponseWriter, r *http.Request) {
	db := core.DbConnMySQL()
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string

		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Show", emp)

	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := core.DbConnMySQL()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Edit", emp)

	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := core.DbConnMySQL()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := core.DbConnMySQL()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := core.DbConnMySQL()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(emp)

	log.Println("DELETE")

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}
