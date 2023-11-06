package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const SERVERURL = "http://localhost:8080"
const OUTPUTDIR = "output/"
const RECEIVERADDR = "http://localhost:8080"

type DownResp struct {
	ip       string
	filename string
	ok       bool
}

type UpResp struct {
	key string
	ok  bool
}

func up(path *string) error {
	resp, err := http.Post(
		SERVERURL,
		"text/plain",
		// Grabs filename from the end of the path
		bytes.NewBuffer(([]byte(filepath.Base(*path)))),
	)
	if err != nil {
		return err
	}

	var body UpResp
	json.NewDecoder(resp.Body).Decode(&body)
	if !body.ok {
		return fmt.Errorf("server error. sorry !")
	}

	fmt.Println("*** success. key is " + body.key + " ***")

	// listen for requests
	listener, err := net.Listen("tcp", RECEIVERADDR)
	if err != nil {
		return err
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	file, err := os.Open(*path)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, conn)
	if err != nil {
		return err
	}

	fmt.Println("Do you want to remove this file? (y/N)")
	var ans string
	fmt.Scanln(&ans)
	if ans == "y" || ans == "Y" {
		err = os.Remove(*path)
		if err != nil {
			return err
		}
	}

	return nil
}

func down(key *string) error {
	// Request server with key
	resp, err := http.Get(SERVERURL + "?key=" + *key)
	if err != nil {
		return err
	}

	// Server responds with IP of target
	var body DownResp
	json.NewDecoder(resp.Body).Decode(&body)
	if !body.ok {
		return fmt.Errorf("invalid key")
	}

	/*
		resp, err = http.Get(body.ip)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	*/
	// THIS OR
	/*
		conn, err := net.Dial("tcp", body.ip)
		if err != nil {
			return err
		}
		defer conn.Close()
	*/

	out, err := os.Create(OUTPUTDIR + body.filename)
	if err != nil {
		return err
	}

	// Download file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("success. saved file to path " + "output")
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("USAGE: sink COMMAND PATH/KEY")
		return
	}

	key := os.Args[2]
	var err error

	switch os.Args[1] {
	case "up":
		var dir string
		dir, err = os.Getwd()
		if err != nil {
			break
		}

		path := filepath.Join(dir, key)
		_, err = os.Stat(path)
		if err != nil {
			break
		}
		err = up(&path)
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
