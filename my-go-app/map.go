package main

import "fmt"

func maps() {
	person := map[string]string{
		"name": "Imam",
		"age":  "25",
	}

	fmt.Println(person["name"]) // Output: Imam

	// Tambah key baru
	person["job"] = "Data Engineer"
	fmt.Println(person)
}

// func main() {
// 	maps()
// }
