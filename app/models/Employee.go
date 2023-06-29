package models

import (
	"main/core"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func GetAll() {
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
