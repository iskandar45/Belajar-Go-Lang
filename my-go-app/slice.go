package main

import "fmt"

func slice() {
	angka := []int{10, 20, 30} // Slice
	angka = append(angka, 40)  // Tambah elemen ke slice
	fmt.Println(angka)         // Output: [10 20 30 40]
}

// func main() {
// 	slice()
// }
