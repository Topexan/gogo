package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/* resp, err := http.Get("http://localhost:8000/persons")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(body)
	log.Println(string(body)) */
	requestBody, err := json.Marshal(map[string]string{
		"id":      "4",
		"name":    "Name4",
		"surname": "Surname4",
		"city":    "City4",
		"phone":   "8747",
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8000/persons", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
