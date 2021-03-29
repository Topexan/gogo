package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if (len(args) == 4) {
		a, errA := strconv.Atoi(args[1])
		b, errB := strconv.Atoi(args[3])
		if errA != nil {
			fmt.Println(errA)
		} else if errB != nil {
			fmt.Println(errB)
		}
		if (args[2] == "+") {
			fmt.Println( a, "+", b, "=", a + b)
		} else if (args[2] == "-") {
			fmt.Println( a, "-", b, "=", a - b)
		} else if (args[2] == "*") {
			fmt.Println( a, "*", b, "=", a * b)
		} else if (args[2] == "/") {
			if (b != 0) {
				fmt.Println( a, "/", b, "=", a / b)
			} else {
				fmt.Println("No division by 0")
			}
		} else if (args[2] == "%") {
			if (b != 0) {
				fmt.Println( a,"%",b,"=",a % b)
			} else {
				fmt.Println("No modulo by 0")
			}
		}
	} else {
		fmt.Println("")
	}
}