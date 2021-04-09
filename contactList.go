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
	json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var p Person
	_ = json.NewDecoder(r.Body).Decode(&p)
	p.ID = strconv.Itoa(rand.Intn(20))
	p.Name = "Sergey"
	p.Surname = "Orlov"
	p.City = "Deputatskiy"
	p.Phone = "unknown"
	fmt.Println(params)
	persons = append(persons, p)
	json.NewEncoder(w).Encode(p)

	db := setupDB()
	defer db.Close()

	sqlStatement := `
		INSERT INTO persons (id, name, surname, city, phone)
		VALUES ($1, $2, $3, $4, $5)`
		_, err := db.Exec(sqlStatement, p.ID, p.Name, p.Surname, p.City, p.Phone)
		if err != nil {
			panic(err)
		}
		fmt.Println("insertion comleted")
}

func getName(w http.ResponseWriter, r *http.Request) 

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	/* keys, ok := r.URL.Query()["name"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]

    fmt.Println("Url Param 'key' is: " + string(key)) */

	params := mux.Vars(r)
	fmt.Println(params)
	var pNew Person
	pNew.ID = params["id"]
	pNew.Name = r.FormValue("name")//params["name"]//"NewName"
	pNew.Surname = r.FormValue("surname")//params["surname"]//"NewSurname"
	pNew.City = r.FormValue("city")//params["city"]//"NewCity"
	pNew.Phone = r.FormValue("phone")//params["phone"]//"NewPhone"
	fmt.Println(pNew.ID)
	fmt.Println(pNew.Name)
	fmt.Println(pNew.Surname)
	fmt.Println(pNew.City)
	fmt.Println(pNew.Phone)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons[index].Name = pNew.Name
			persons[index].Surname = pNew.Surname
			persons[index].City = pNew.City
			persons[index].Phone = pNew.Phone
			json.NewEncoder(w).Encode(persons[index])
			return
		}
	}
	json.NewEncoder(w).Encode(persons)
/* 
	db := setupDB()
	defer db.Close()

	sqlStatement := `
		UPDATE persons SET name = $2,
						   surname = $3,
						   city = $4,
						   phone = $5
		WHERE id = $1;`
		_, err := db.Exec(sqlStatement, params["id"], pNew.Name, pNew.Surname, pNew.City, pNew.Phone)
		if err != nil {
			panic(err)
		} */

}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)

	db := setupDB()
	defer db.Close()

	sqlStatement := `
		DELETE FROM persons
		WHERE name = $1;`
		_, err := db.Exec(sqlStatement, params["id"])
		if err != nil {
			panic(err)
		}
}

func connectToDb() {
	db := setupDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

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
}

func setupDB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	connectToDb()
	r := mux.NewRouter()
	r.HandleFunc("/persons", getPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", getPerson).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons/", updatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id}", deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
