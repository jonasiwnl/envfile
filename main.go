package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jonasiwnl/mm/mongo"
	"github.com/jonasiwnl/mm/parser"
)

func main() {
	log.Println("starting...")

	env, err := ValidateEnv()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("asdf") // TODO
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Start client
	client, err := mongo.NewClient(env.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Lex migration file
	parser.NewLexer(file).Lex()

	// Parse migration file
	parser.NewParser().Parse()

	// Make database changes

	// Clean up
	os.Remove("temp") // TODO
	client.Disconnect(ctx)

	log.Println("done.")
}
