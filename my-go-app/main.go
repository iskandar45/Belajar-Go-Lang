package main

import "fmt"

func main() {
	fmt.Println("Hello, GitHub Codespaces!")
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("1. Variable & Type Data")
	type_data()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("2. Function & Loop")
	loop()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("3. Array & Slice")
	array()
	slice()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("4. Map (Mirip Object di JavaScript atau Dictionary di Python)")
	maps()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("5. (Mirip Class di OOP, Tapi Tanpa Method)")
	structs()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("6. Pointer (Konsep Memory Golang)")
	pointers()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("7. Conditions")
	conditions()
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println("8. Looping")
	looping()
	fmt.Println("")
	fmt.Println("++++++++++++++++++++++++++")
}
