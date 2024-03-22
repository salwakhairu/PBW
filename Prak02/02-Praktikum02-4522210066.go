package main

import (
	"fmt"
)

func main() {
	// Menyimpan data mahasiswa dalam map
	mahasiswaMap := map[string]map[string]string{
		"4522210066": {"Nama": "Salwa Khairu Mista", "Jurusan": "Teknik Informatika"},
		"4322210067": {"Nama": "Farah Tri Mahardini", "Jurusan": "Teknik Arsitek"},
		"4122210068": {"Nama": "Nadia Ayu Rahmawati", "Jurusan": "Teknik Mesin"},
	}

	// Menampilkan daftar nama mahasiswa
	fmt.Println("Daftar Nama Mahasiswa:")
	for _, data := range mahasiswaMap {
		fmt.Println(data["Nama"])
	}

	// Input NPM untuk mencari
	var npmCari string
	fmt.Print("\nMasukkan NPM untuk mencari data mahasiswa: ")
	fmt.Scanln(&npmCari)

	// Mencari data mahasiswa berdasarkan NPM
	if mahasiswa, found := mahasiswaMap[npmCari]; found {
		fmt.Printf("\nData Mahasiswa dengan NPM %s ditemukan:\n", npmCari)
		fmt.Printf("Nama: %s\nNPM: %s\nJurusan: %s\n", mahasiswa["Nama"], npmCari, mahasiswa["Jurusan"])
	} else {
		fmt.Printf("\nData Mahasiswa dengan NPM %s tidak ditemukan.\n", npmCari)
	}
}
