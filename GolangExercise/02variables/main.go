package main

import "fmt"

func main() {
	var userName string = "thambi"
	fmt.Println(userName)
	fmt.Printf("Variable type: %T \n", userName)

	var isTrue bool = false
	fmt.Println(isTrue)
	fmt.Printf("Variable type: %T \n", isTrue)

	var MagicNumber int16 = 256
	fmt.Println(MagicNumber)
	fmt.Printf("Variable type: %T \n", MagicNumber)

	var empty string
	fmt.Println(empty)
	fmt.Printf("Variable type: %T \n", empty)

	// implicit type

	// no var style
	notOk := false
	fmt.Println(notOk)
}
