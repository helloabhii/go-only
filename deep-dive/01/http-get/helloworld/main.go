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
	args := os.Args                                                                   //arguments
	fmt.Printf("helloabhii \n os  Arguments :  %v\n Argument: %v\n", args, args[1:3]) // you  have to give the arguments otherwise cause an error
	// %v means compiler auto decide what the type of that agrument
	//printf because of formatting
	// go run main.go arg1 arg2
}
