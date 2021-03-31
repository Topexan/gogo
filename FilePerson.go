package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args
	checkList := false
	for i, _ := range args {
		if (args[i] == "-list") {
			checkList = true
			break
		}
	}
	if (checkList) {
		data, err := ioutil.ReadFile("PersonList.txt")
		if (err != nil) {
			fmt.Println(err)
		}
		arr := strings.Split(string(data), "\n")
		var arrRes []string
		for i := 1; i < len(arr); i++ {
			arrRes = strings.Split(arr[i], " ")
			for j := 1; j <= len(arrRes); j++ {
				if (j % 4 == 1) {
					fmt.Print("Name: ",arrRes[j-1] )
				} else if (j % 4 == 2) {
					fmt.Print(" Surname: ",arrRes[j-1] )
				} else if (j % 4 == 3) {
					fmt.Print(" City: ", arrRes[j-1] )
				} else if (j % 4 == 0) {
					fmt.Print(" Number: ", arrRes[j-1] )
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("")
	}
}
