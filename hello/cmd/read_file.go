package main

import (
	"fmt"
	"log"
	"os"
)

func try01() {
	body, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
}
