package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const SERVERURL = "http://localhost:3000/"
const OUTPUTDIR = "output/"
const RECEIVERPORT = "8080"

/*
type DownResp struct {
	ip       string
	filename string
	ok       bool
}

type UpResp struct {
	key string
	ok  bool
}
*/

type SendResp struct {
	ip       string
	filename string
	ok       bool
}

type ExpectResp struct {
	key string
	ok  bool
}

// **Endpoints**

// /expect
// Server stores IP and responds with keyword

// /send?key=KEY
// Server responds with IP

func send(path string, keyword string) error {
	// Request server with key, get back IP
	var body SendResp
	resp, err := http.Get(SERVERURL + "send?key=" + keyword + "&filename=" + filepath.Base(path))
	if err != nil {
		return err
	}

	json.NewDecoder(resp.Body).Decode(&body)
	if !body.ok {
		return fmt.Errorf("invalid key (or server error)")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// Create a form field for the filename
	filenameField, err := writer.CreateFormField("filename")
	if err != nil {
		return err
	}
	filenameField.Write([]byte(filepath.Base(path)))

	// Create a form field for the file
	fileField, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return err
	}

	// Copy the file content to the form field
	_, err = io.Copy(fileField, file)
	if err != nil {
		return err
	}
	writer.Close()

	// Create a POST request with `buf` (mulipart form data)
	_, err = http.NewRequest("POST", body.ip, buf)
	if err != nil {
		return err
	}

	fmt.Println("success. sent file to " + body.ip)
	return nil
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	out, err := os.Create(OUTPUTDIR + handler.Filename)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, r.Body)
	if err != nil {
		fmt.Println("Error copying the file:", err)
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return
	}

	fmt.Println("success. recieved file.")
}

func expect(path string) error {
	_, err := os.Stat(path)

	// Does the file exist
	if !os.IsNotExist(err) {
		if err != nil {
			return err
		}

		fmt.Println("this file already exists. do you want to overwrite? (y/N)")
		var ans string
		fmt.Scanln(&ans)
		if ans != "y" && ans != "Y" {
			return nil
		}
	}

	resp, err := http.Get(SERVERURL + "expect?filename=" + path)
	if err != nil {
		return err
	}

	var body ExpectResp
	json.NewDecoder(resp.Body).Decode(&body)

	// Receive keyword from server and print
	fmt.Println("your keyword is...: " + body.key)

	http.HandleFunc("/file", handleFile)
	http.ListenAndServe(":"+RECEIVERPORT, nil)

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: sink COMMAND")
		return
	}

	var err error

	switch os.Args[1] {
	case "expect":
		if len(os.Args) < 3 {
			fmt.Println("USAGE: sink expect OUTPUT_PATH")
			return
		}
		err = expect(os.Args[2])
	case "send":
		if len(os.Args) < 4 {
			fmt.Println("USAGE: sink send FILE_PATH KEYWORD")
			return
		}
		err = send(os.Args[2], os.Args[3])
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

/*
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
	case "expect":
		err = expect(&key)
	case "send":
		err = send(filepath.Base(keyword), &key)
	default:
		fmt.Println("unknown command.")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("exiting.")
}*/

/*
func up(path *string) error {
	// TODO merge this with line 49
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

	client := sse.NewClient(serverURL)
	events := client.Subscribe()

	for {
		select {
		case event := <-events:
			// Split the event data by newlines
			dataLines := strings.Split(event.Data, "\n")

			if len(dataLines) >= 2 {
				// Extract the key and IP address
				key := dataLines[0]
				ipAddress := dataLines[1]

				// Print the key
				fmt.Println("Received key:", key)

				// Send a POST request to the IP address
				err := sendPOSTRequest(ipAddress)
				if err != nil {
					fmt.Println("Error sending POST request:", err)
				}
			}
		}
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
/*

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
*/
