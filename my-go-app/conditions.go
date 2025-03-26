package main

import "fmt"

func basicCondition() {
	umur := 25

	if umur > 18 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Belum Dewasa")
	}
}

func multipleCondition() {
	nilai := 75

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 75 {
		fmt.Println("B")
	} else if nilai >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

func shortStatement() {
	// umur hanya bisa diakses di dalam blok if-else.
	if umur := 20; umur >= 18 {
		fmt.Println("Boleh Buat SIM")
	} else {
		fmt.Println("Belum Boleh Buat SIM")
	}
}

func switchCaseStatement() {
	hari := "Senin"

	switch hari {
	case "Senin":
		fmt.Println("Hari Kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur")
	default:
		fmt.Println("Hari Biasa")
	}
}

func conditions() {
	basicCondition()
	multipleCondition()
	shortStatement()
	switchCaseStatement()
}
