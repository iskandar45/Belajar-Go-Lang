package main

import "fmt"

func basicLooping() {
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
}

func loopWithhoutInitPostStatement() {
	angka := 5
	for angka <= 5 {
		fmt.Println("Perulangan ke -", angka)
		angka++
	}
}

func loopForRangeStatement() {
	// looping yang biasa digunakan untuk mengulang elemen dari array, slice, map, atau channel
	angka := []int{10, 20, 30}
	for i, v := range angka {
		fmt.Println("Index: ", i, " Value: ", v)
	}
	for _, v := range angka {
		fmt.Println("Value: ", v)
	}
}

func loopingMap() {
	person := map[string]string{
		"name": "Nur Imam Iskandar",
		"job":  "Data Engineer",
	}
	//For-range di map selalu mengembalikan key, value.
	for key, value := range person {
		fmt.Print("\n", key, " : ", value)
	}
}

func looping() {
	basicLooping()
	loopWithhoutInitPostStatement()
	loopForRangeStatement()
	loopingMap()
}
