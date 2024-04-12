package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	url := "http://www.google.com"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
