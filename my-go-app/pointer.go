package main

import "fmt"

func ubahUmur(umur *int) {
	*umur = 27
}

func pointers() {
	age := 25
	fmt.Println("Before Umur: ", age) // Output: 25
	ubahUmur(&age)
	fmt.Println("After Umur: ", age) // Output: 27
}

// func main() {
// 	pointers()
// }
