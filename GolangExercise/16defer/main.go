package main

import "fmt"

func main() {
	defer fmt.Println("zero")
	fmt.Println("one")
	count()

}

func count() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
