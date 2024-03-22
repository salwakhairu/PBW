package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Masukkan usia: ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	// KATEGORI USIA
	kategori := ""
	if usia < 18 {
		kategori = "Anak-anak"
	} else {
		kategori = "Dewasa"
	}

	// MENAMPILKAN PESAN
	fmt.Printf("Selamat datang %s, Anda termasuk kategori %s.\n", nama, kategori)
}
