package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	// Mengubah izin direktori.
	err = os.Chmod("salwakhairumista", 0066)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Izin 'salwakhairumista' telah diubah menjadi 0066")
}
