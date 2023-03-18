package main

import (
	"golang_challenge/pkg/db"
	"golang_challenge/pkg/handles"
	"log"
)

func main() {
	// This function will make a connection pool to use the database.
	err := db.InitDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// This fuction its going to create a new table in the database.
	db.Migration()

	// This function makes a request to https://jsonplaceholder.typicode.com/users API
	// and will return an array and a nil in case that there isn' an error.
	users, err := handles.GetUsers()

	// This function is used to save the result we got from handles.GetUsers().
	err = db.PopulateUsersTable(users)
	if err != nil {
		log.Fatalln(err)
	}

}
