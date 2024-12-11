package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	//since i only need to check the err value and dont need the response value
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println("This URL is invalid : ", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("This is the body %s and this is the statuscode: %d ", body, response.StatusCode)
}
