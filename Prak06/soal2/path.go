package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	filelnfo, err := os.Stat("salwakhairumista")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if filelnfo.IsDir() {
		fmt.Println("salwakhairumista adalah sebuah direktori.")
	} else {
		fmt.Println("salwakhairumista adalah sebuah file.")
	}
}
