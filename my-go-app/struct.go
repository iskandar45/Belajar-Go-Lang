package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func structs() {
	var imam Person = Person{Name: "Imam", Age: 25}
	fmt.Println(imam.Name) // Output: Imam
}

// func main() {
// 	structs()
// }
