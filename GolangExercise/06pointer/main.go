package main

import "fmt"

func main() {

	myNumber := 2

	var ptr = &myNumber

	fmt.Println("the ptr value :", ptr)

	fmt.Println("the value  at the address :", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value of MyNumber : ", myNumber)
}
