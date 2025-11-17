package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["backend"] = "dev"
	// fmt.Println(m["name"], m["backend"])

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 30
	// // fmt.Println(m["age"])
	// // fmt.Println(len(m))
	// delete(m, "price")
	// fmt.Println(m)
	// m := map[string]int{"price": 40, "phones": 20}
	// // fmt.Println(m)
	// v, ok := m["price"]
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	// fmt.Println(v)
	m1 := map[string]int{"price": 40, "phones": 20}
	m2 := map[string]int{"price": 40, "phones": 20}
	fmt.Println(maps.Equal(m1, m2))

}
