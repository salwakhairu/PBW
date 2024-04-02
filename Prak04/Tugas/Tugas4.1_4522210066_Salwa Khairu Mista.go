package main

import "fmt"

func fact(jumlah int) int {
	if jumlah == 1 {
		return 1
	}
	return jumlah * fact(jumlah-1)
}

func main() {
	// Fungsi dibawah sama saja dengan
	// 7 x (6 x (5 x (4 x (3 x (2 x (1 x 1))))))
	fmt.Println(fact(7))
}
