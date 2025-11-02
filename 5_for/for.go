package main

import (
	"fmt"
)

func main() {
	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	//infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	//classic for loop

	// for i := 0; i < 5; i++ {
	// 	// break
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// go v 1.22
	for i := range 5 {
		fmt.Println(i)
	}

}
