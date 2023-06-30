package core

import (
	"context"
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnMySQL() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := "0mEg4a9012_"
	dbName := "crud-db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func DbConnMongo() *mongo.Client {
	// Define the mongodb client URL
	var uri = "mongodb://localhost:27017"

	// Establish the connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Create go routine to defer the closure
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin insertOne and create testDB database
	coll := client.Database("testDB").Collection("scoreCollection")
	doc := bson.D{
		{"name", "Anna"},
		{"score", 9.5},
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	// end insertOne

	// When you run this file, it should print:
	// Document inserted with ID: ObjectID("...")
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)

	return client
}
