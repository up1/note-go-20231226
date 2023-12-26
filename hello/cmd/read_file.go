package main

import (
	"fmt"
	"iotuil"
	"log"
)

func main() {
	body, err := iotuil.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
}
