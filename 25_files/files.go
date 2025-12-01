// write a go code that creates a file inside current dirrectory names hello.txt
package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a file named "hello.txt" in the current directory
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File 'hello.txt' created successfully")
}
