package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func loop() {
	salam := greet("Nur Imam")
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i, " - ", salam)
	}
}

// func main() {
// 	loop()
// }
