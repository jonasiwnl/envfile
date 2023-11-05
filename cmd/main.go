package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO figure out how to do p2p

func up(path *string) error {
	// TODO read file and POST req
  fmt.Println("Not implemented yet!")
	return nil
}

func down(key *string) error {
	fmt.Println(*key)

	// TODO link
	resp, err := http.Get("https://jonasiwnl.github.io/Jonas_Groening_resume.pdf")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("output") // TODO how to get filename?
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

  fmt.Println("success. saved file to path " + "output")
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("USAGE: sink COMMAND PATH/KEY")
		return
	}

	key := os.Args[1]
	var err error

	switch os.Args[1] {
	case "up":
		err = up(&key)
	case "down":
		err = down(&key)
	default:
		fmt.Println("unknown command.")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("exiting.")
}
