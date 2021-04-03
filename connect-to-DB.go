package main

import (
  "database/sql"
  "fmt"
  "os"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "enteam"
  dbname   = "go_practice"
)

type Person struct {
	id        string
	name      string
	surname   string
	city      string
	phone     string
  }
  

func main() {
	checkInsert := false
	checkSelect := false
	checkList := false
	checkDelete := false
	args := os.Args
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Successfully connected!")
	for i := range args {
		if (args[i] == "-insert") {
			checkInsert = true
		}
		if (args[i] == "-select") {
			checkSelect = true
		}
		if (args[i] == "-list") {
			checkList = true
		}
		if (args[i] == "-delete") {
			checkDelete = true
		}
	}

	if (checkInsert) {
		sqlStatement := `
		INSERT INTO persons (id, name, surname, city, phone)
		VALUES ($1, $2, $3, $4, $5)`
		_, err = db.Exec(sqlStatement, args[2], args[3], args[4], args[5], args[6])
		if err != nil {
			panic(err)
		}
		fmt.Println("insertion comleted")
	}

	if (checkSelect) {
		sqlStatement := `SELECT id, name FROM persons WHERE id=$1;`
		var name string
		var id string
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := db.QueryRow(sqlStatement, args[2])
		switch err := row.Scan(&id, &name); err {
			case sql.ErrNoRows:
				fmt.Println("No rows were returned!")
			case nil:
				fmt.Println(id, name)
			default:
				panic(err)
		}
	}

	if (checkList) {
		rows, err := db.Query("SELECT id, name, surname, city, phone FROM persons")
		if err != nil {
			// handle this error better than this
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id string
			var name string
			var surname string
			var city string
			var phone string
			err = rows.Scan(&id, &name, &surname, &city, &phone)
			if err != nil {
				// handle this error
				panic(err)
			}
			fmt.Println(id, name, surname, city, phone)
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			panic(err)
		}
	}

	if (checkDelete) {
		sqlStatement := `
		DELETE FROM persons
		WHERE id = $1;`
		_, err = db.Exec(sqlStatement, args[2])
		if err != nil {
			panic(err)
		}
	}
}