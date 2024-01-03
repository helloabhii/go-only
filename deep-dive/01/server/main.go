package main

import (
	"fmt"
	"net/http"
)

func main() {

	server()
}

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "helloabhii")
}
func server() {
	http.HandleFunc("/", handler)
	port := 8080
	fmt.Printf("Server listening on : %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("ErrorL %s\n", err)
	}
}
