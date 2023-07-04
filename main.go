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
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Start client
	client, err := mongo.NewClient(env.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Lex & parse migration file
	parser.NewLexer(file).Lex()
	parser.Parse()

	// Make database changes
	mongo.Execute(ctx, client)

	// Clean up
	os.Remove("temp") // TODO
	client.Disconnect(ctx)

	log.Println("done.")
}
