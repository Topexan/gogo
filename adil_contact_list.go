package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "qwerty"
	DB_NAME     = "contact_list"
)

type Person struct {
	PersonID int `json:"personid"`
	PersonName string `json:"personname"`
	PersonSurname string `json:"personsurname"`
	PersonCity string `json:"personcity"`
	PersonPhone string `json:"personphone"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Person `json:"data"`
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	// Get all persons
	router.HandleFunc("/persons/", GetPersons).Methods("GET")

	// Create a person
	router.HandleFunc("/persons/", CreatePerson).Methods("POST")

	// Delete a specific person by the personID
	router.HandleFunc("/persons/{personid}", DeletePerson).Methods("DELETE")

	// Delete all persons
	router.HandleFunc("/persons/", DeletePersons).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get all books
func GetPersons(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting persons...")

	// Get all books from books table that don't have personID = "1"
	rows, err := db.Query("SELECT * FROM persons where id <> $1", "1")

	checkErr(err)
	var persons []Person
	// var response []JsonResponse
	// Foreach book
	for rows.Next() {
		var id int
		var name string
		var surname string
		var city string
		var phone string

		err = rows.Scan(&id, &name, &surname, &city, &phone)

		checkErr(err)

		persons = append(persons, Person{PersonID: id, PersonName: name, PersonSurname: surname, PersonCity: city, PersonPhone: phone})
	}

	var response = JsonResponse{Type: "success", Data: persons}

	json.NewEncoder(w).Encode(response)
}

// Create a book
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	//name := r.FormValue("name")
	//surname := r.FormValue("surname")
	//phone := r.FormValue("phone")
	//city := r.FormValue("city")

	decoder := json.NewDecoder(r.Body)
	var p Person
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	var response = JsonResponse{}

	if p.PersonName == "" || p.PersonSurname == "" || p.PersonPhone == "" || p.PersonCity == "" {
		response = JsonResponse{Type: "error", Message: "You are missing name or surname or phone or city parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting person into DB")

		fmt.Println("Inserting new person with name: " + p.PersonName + " and surname: " + p.PersonSurname + " and city: " + p.PersonCity + " and phone: " + p.PersonPhone)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO persons(name, surname, city, phone) VALUES($1, $2, $3, $4) returning id;", p.PersonName, p.PersonSurname, p.PersonCity, p.PersonPhone).Scan(&lastInsertID)

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The person has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a book
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	personID := params["personid"]

	var response = JsonResponse{}

	if personID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing personID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting person from DB")

		_, err := db.Exec("DELETE FROM persons where id = $1", personID)
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The person has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all books
func DeletePersons(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all persons...")

	_, err := db.Exec("DELETE FROM persons")
	checkErr(err)

	printMessage("All persons have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All persons have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}