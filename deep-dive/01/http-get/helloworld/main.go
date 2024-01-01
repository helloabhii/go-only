package main

import (
	"fmt"
	"os"
)

//Executable always in the mainfile

//go build main.go - this will make a executable file inside the folder

func main() {
	argu()

}

func argu() {
	args := os.Args //arguments
	fmt.Printf("helloabhii \n Arguments :  %v\n", args)
	// %v means compiler auto decide what the type of that agrument
	//printf because of formatting
}
