package envfile

import (
	"log"
	"os"
)

type Environment struct {
	Vars map[string]string
}

func SetUp() (func(), error) {
	log.Println("ENVFILE: starting...")

	file, err := os.ReadFile("Envfile")
	if err != nil {
		return nil, err
	}

	// Lex & parse Envfile
	NewLexer(string(file)).Lex()
	env, err := Parse()
	if err != nil {
		return nil, err
	}

	// Set environment variables
	// for _, var := range env.Vars {
	// 	os.Setenv(var.Name, var.Value)
	// }

	return env.TearDown, nil
}

func (e *Environment) TearDown() {
	log.Println("ENVFILE: cleaning up...")

	log.Println("ENVFILE: done.")
}
