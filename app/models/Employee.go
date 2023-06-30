package models

import (
	"log"
	"main/core"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func GetAll() []Employee {
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

	return res
}

func GetById(nId string) Employee {
	db := core.DbConnMySQL()

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

	defer db.Close()

	return emp
}

func Create(name string, city string) {
	db := core.DbConnMySQL()
	insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}

	insForm.Exec(name, city)
	log.Println("INSERT: Name: " + name + " | City: " + city)
	defer db.Close()
}

func UpdateById(name string, city string, id string) {
	db := core.DbConnMySQL()
	insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	insForm.Exec(name, city, id)
	log.Println("UPDATE: Name: " + name + " | City: " + city)
	defer db.Close()
}

func DeleteById(id string) {
	db := core.DbConnMySQL()
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(id)

	log.Println("DELETE")

	defer db.Close()
}
