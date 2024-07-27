package main

import "fmt"

func myFunction(i int) int {
	return i % 2
}

func main() {
	fmt.Println(myFunction(100))
}
