package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
	Phone   string `json:"phone"`
}

//function to show all persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var persons []Person
	db := setupDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

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
		persons = append(persons, Person{ID: id, Name: name, Surname: surname, City: city, Phone: phone})
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	log.Println("Get all persons from table: persons")
	json.NewEncoder(w).Encode(persons)
}

//function to show one person
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //read parameter from url

	db := setupDB()
	defer db.Close()

	sqlStatement := `
		SELECT id, name, surname, city, phone FROM persons
		WHERE id = $1;`
	rows, err := db.Query(sqlStatement, params["id"])
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var p Person
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
		p = Person{ID: id, Name: name, Surname: surname, City: city, Phone: phone}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	log.Println("Get person from table: persons, with id:", params["id"])
	json.NewEncoder(w).Encode(p)
}

//function to create new person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var p Person //new Person
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	_ = json.NewDecoder(r.Body).Decode(&p)

	p.ID = strconv.Itoa(rand.Intn(30-4) + 4) // random number from 4 to 30
	json.NewEncoder(w).Encode(p)

	db := setupDB()
	defer db.Close()

	sqlStatement := `
		INSERT INTO persons (id, name, surname, city, phone)
		VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(sqlStatement, p.ID, p.Name, p.Surname, p.City, p.Phone) //add person to database
	if err != nil {
		panic(err)
	}
	log.Println("Insert new person to table: persons")
}

// function to update existing person
func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var p Person //person with new parameters
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(p)

	db := setupDB()
	defer db.Close()

	sqlStatement := `
		UPDATE persons SET name = $2,
						   surname = $3,
						   city = $4,
						   phone = $5
		WHERE id = $1;`
	_, err = db.Exec(sqlStatement, p.ID, p.Name, p.Surname, p.City, p.Phone) //update person in our database
	if err != nil {
		panic(err)
	}
	log.Println("Update person from table: persons, with id:", p.ID)

}

//function to delete person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //take id from url

	db := setupDB()
	defer db.Close()

	sqlStatement := `
		DELETE FROM persons
		WHERE id = $1;`
	_, err := db.Exec(sqlStatement, params["id"]) //delete person from database
	if err != nil {
		panic(err)
	}
	log.Println("Delete person from table: persons, with id:", params["id"])
	getPersons(w, r)
}

//function to connect to database
func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//panic(err)
		log.Println(err)
	}
	log.Println("Succesful connetcion to ", dbname)
	return db
}

func main() {
	r := mux.NewRouter() //create new router
	r.HandleFunc("/persons", getPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", getPerson).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons", updatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id}", deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
