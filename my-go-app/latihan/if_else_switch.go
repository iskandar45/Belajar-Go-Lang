package main

import (
	"fmt"
	"os"
)

func ifElseSwitch() {
	var umur int

	fmt.Print("Masukkan Umur Anda: ")
	_, err := fmt.Scanln(&umur)

	if err != nil {
		fmt.Println("❌ Input tidak valid! Harus berupa angka.")
		os.Exit(1) // keluar dari program
	}

	if umur < 0 {
		fmt.Println("❌ Umur tidak boleh negatif.")
		os.Exit(1)
	}

	switch {
	case umur < 12:
		fmt.Println("👶 Kamu masih anak-anak")
	case umur >= 12 && umur <= 17:
		fmt.Println("🧒 Kamu remaja")
	case umur >= 18 && umur <= 25:
		fmt.Println("🧑 Kamu dewasa muda")
	case umur > 25:
		fmt.Println("🧔 Kamu dewasa")
	default:
		fmt.Println("⚠️ Input tidak dikenali")
	}
}
