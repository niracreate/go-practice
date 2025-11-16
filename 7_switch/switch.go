package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its work day")
	}
	//type switch
	// whoAmI := func(i any) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("Its an integer")
	// 	case string:
	// 		fmt.Println("Its an string")
	// 	case bool:
	// 		fmt.Println("Its an boolean")
	// 	default:
	// 		fmt.Println("other", t)
	// 	}
	// }
	// whoAmI(true)

}
