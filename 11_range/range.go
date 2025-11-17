package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8}
	// sum := 0

	// for _, num := range nums {
	// 	sum = sum + num
	// }

	// fmt.Println(sum)
	// m := map[string]string{"fname": "john", "lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
