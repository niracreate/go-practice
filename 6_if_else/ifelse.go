package main

import "fmt"

func main() {
	// age := 18
	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// if age := 18; age >= 18 {
	// 	fmt.Println("Adult")

	// } else if age >= 12 {
	// 	fmt.Println("Teenager")
	// } else {
	// 	fmt.Println("Kid")
	// }
	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	//go doesnot have ternary operator

}
