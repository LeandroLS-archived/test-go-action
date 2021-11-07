package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Olá")
	CreateFile()
}

func CreateFile() {

	_, err := os.Create("aaaa.txt")
	if err != nil {
		fmt.Println(err)
	}
	// return file
}
