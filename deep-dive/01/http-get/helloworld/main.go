package main

import (
	"fmt"
	"os"
)

//Executable always in the mainfile

//go build main.go - this will make a executable file inside the folder

func main() {
	argu()
	fund()

}

func argu() {
	args := os.Args

	// if len(args) < 2 {
	// 	fmt.Printf("Usage: ./ hello-world <argument>\n")
	// 	os.Exit(1) // means unsuccessful
	// }
	//arguments
	fmt.Printf("helloabhii \n os  Arguments :  %v\n Argument: %v\n", args, args[1:]) // you  have to give the arguments otherwise cause an error
	// %v means compiler auto decide what the type of that agrument
	//printf because of formatting
	// go run main.go arg1 arg2
}

//slice and array
// slice can be of dynamic length
// arrays are of static length

//to check program successfully executed or not -
// echo $?

/// map ->

// var  a map[string]string    // here [string] = key type , string = value type
// a  := make(map[string]string)  // short hand to create map
// a := map[string]string{"key":"value"}

// array
// var arr [2]int = [2]int{1,2}   // array includes [1,2]
// array of size 2 with integer type

// slice
// var slc []string // declaring slice with type string and size is not mentioned so it is dynamic by default
// a := make([]string,5) //size 5
// a := make([]string,5,10) // size 5 with capacity 10

// declaring multiple variables

func fund() {
	var (
		a int
		b string
	)
	a = 69
	b = "abhi"
	fmt.Print(b)
	fmt.Println(a)
}
