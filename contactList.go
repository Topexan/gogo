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
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "enteam"
	dbname   = "go_practice"
)

type Person struct {
	ID      string `json:"id"`
	Name    string `json:"title"`
	Surname string `json:"author"`
	City    string `json:"city"`
	Phone   string `json:"phone"`
}

var persons []Person

func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//fmt.Println(params)
	json.NewEncoder(w).Encode(&Person{})
}

func connectToDb() {
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

	rows, err := db.Query("SELECT id, name, surname, city, phone FROM persons")
	if err != nil {
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
			panic(err)
		}
		persons = append(persons, Person{ID: id, Name: name, Surname: surname, City: city, Phone: phone})
		fmt.Println(id, name, surname, city, phone)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	connectToDb()
	r := mux.NewRouter()
	r.HandleFunc("/persons", getPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", getPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
