package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// {"page":"words", "words": ["word1"]}

// struct
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"` // array of string
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format : %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1]) // golang has http pacakge in standard library
	//response have status code

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Raw Response:", string(body))

	//only parse data if status code is 200
	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP code %d): %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words
	err = json.Unmarshal(body, &words)
	if err != nil {
		fmt.Printf("Error Unmarshalling JSON: %s\n", err)
	}

	fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))
}

//parse the output
//parse json into golang variable
