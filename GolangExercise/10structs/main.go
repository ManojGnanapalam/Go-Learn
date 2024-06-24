package main

import "fmt"

func main() {
	// no inheritance in golang ; no super or parent

	naveen := User{"Naveen", "naveen@ok.com", true, 21}
	fmt.Printf("%+v\n", naveen)
	fmt.Printf("%v is  %v year old boy\n", naveen.Name, naveen.Age)
	fmt.Printf("%T", naveen)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
