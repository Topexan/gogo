package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type person struct {
	number      int
	name        string
	surname     string
	city        string
	phoneNumber string
}

func main() {
	args := os.Args
	checkList := false
	checkAdd := false
	checkGet := false
	checkDelete := false

	arrPerson := listPerson("PersonList.txt")

	for i, _ := range args {
		if args[i] == "-list" {
			checkList = true
		}
		if args[i] == "-add" {
			checkAdd = true
		}
		if args[i] == "-get" {
			checkGet = true
		}
		if args[i] == "-delete" {
			checkDelete = true
		}
	}
	if checkList {
		for i := range arrPerson {
			printPerson(arrPerson[i])
		}
	}

	if checkAdd {
		file, err := os.OpenFile("PersonList.txt", os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			fmt.Println(err)
		} else {
			newPerson := strconv.Itoa(len(arrPerson)+1) + " " + args[2] + " " + args[3] + " " + args[4] + " " + args[5]
			n4, err := file.WriteString("\n" + newPerson)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("wrote %d bytes\n", n4)
		}
		file.Close()
	}

	if checkGet {
		for i := range arrPerson {
			if args[2] == strconv.Itoa(arrPerson[i].number) {
				printPerson(arrPerson[i])
			}
		}
	}
	if checkDelete {
		var arrPersonRes []person
		for i := 0; i < len(arrPerson); i++ {
			if strconv.Itoa(arrPerson[i].number) != args[2] {
				arrPersonRes = append(arrPersonRes, arrPerson[i])
			}
		}
		os.Remove("PersonList.txt")
		file, err := os.Create("PersonList.txt")
		file.WriteString("2")
		if err != nil {
			fmt.Println(err)
		} else {
			for i := range arrPersonRes {
				ps := strconv.Itoa(arrPersonRes[i].number) + " " + arrPersonRes[i].name + " " + arrPersonRes[i].surname + " " + arrPersonRes[i].city + " " + arrPersonRes[i].phoneNumber
				n, err := file.WriteString("\n" + ps)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("wrote %d bytes\n", n)
			}
		}
	}
}

func listPerson(s string) []person {
	var arrPerson []person
	data, err := ioutil.ReadFile(s)
	var arrRes []string
	arr := strings.Split(string(data), "\n")
	arrPerson = make([]person, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		arrRes = strings.Split(arr[i], " ")
		arrPerson[i-1].number, err = strconv.Atoi(arrRes[0])
		arrPerson[i-1].name = arrRes[1]
		arrPerson[i-1].surname = arrRes[2]
		arrPerson[i-1].city = arrRes[3]
		arrPerson[i-1].phoneNumber = arrRes[4]
	}

	if err != nil {
		fmt.Println(err)
	}
	return arrPerson
}

func printPerson(person person) {
	fmt.Println(person.number, "Name:", person.name, "Surname:", person.surname, "City:", person.city, "Number:", person.phoneNumber)
}
