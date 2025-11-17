package main

import "fmt"

func changeNum(num *int) { //holds address of an int value
	*num = 5 //deferencing to change or modify the value that pointer points to
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num) //&num gives the address of num.
	// fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)
	// a := 2
	// p := &a
	// fmt.Println(*p)

}
