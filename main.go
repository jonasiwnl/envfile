package main

import (
	"log"
	"os"
)

func main() {
	// Lex conf file
	// Parse conf file

	// Read old conf
	data, err := os.ReadFile("config.mi.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)

	// Compare

	// Start client
	// Make database changes
	// Disconnect client

	log.Println("Done")
}
