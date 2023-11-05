package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

type Resp struct {
	data string
	ok   bool
}

func up(path *string) error {
	resp, err := http.Post(SERVERURL, "text/plain", bytes.NewBuffer(([]byte(*path))))
	if err != nil {
		return err
	}

	var key Resp
	json.NewDecoder(resp.Body).Decode(&key)
	if !key.ok {
		return fmt.Errorf("server error. sorry !")
	}

	fmt.Println("*** success. key is " + key.data + " ***")

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

	_, err = io.Copy(*path, conn)
	if err != nil {
		return err
	}

	fmt.Println("Do you want to remove this file? (y/N)")
	var ans string
	fmt.Scanln(&ans)
	if ans == "y" {
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
	var ip Resp
	json.NewDecoder(resp.Body).Decode(&ip)
	if !ip.ok {
		return fmt.Errorf("invalid key")
	}

	/*
		resp, err = http.Get(ip.data)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	*/
	// THIS OR
	/*
		conn, err := net.Dial("tcp", ip.data)
		if err != nil {
			return err
		}
		defer conn.Close()
	*/

	out, err := os.Create(OUTPUTDIR + "output1") // TODO how to get filename?
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
