package core

import "database/sql"

func DbConnMySQL() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := "password"
	dbName := "crud-db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
