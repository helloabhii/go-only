package main

import (
	"io"
	"log"
	"net/http"
  "fmt"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Helloabhii")
		data, err := io.ReadAll(r.Body)
    if err != nil {
      http.Error(rw, "Oops", http.StatusBadRequest)
      return
    }
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbyeabhii")
	})
	http.ListenAndServe(":9090", nil)
}
