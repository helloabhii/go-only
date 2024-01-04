package main

import (
	"fmt"
	"io"
	"log"
)

// struct
type MySlowReader struct {
	contents string
	pos      int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos <= len(m.contents) {
		n := copy(p, m.contents[m.pos:m.pos+1])
		//  0 <= 1
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {
	mySlowReaderInstance := &MySlowReader{
		contents: "abhi",
	}

	out, err := io.ReadAll(mySlowReaderInstance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s\n", out)
}
