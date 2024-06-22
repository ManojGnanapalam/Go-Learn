package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your wish")

	// comma ok  || err ok

	input, _ := reader.ReadString('\n')
	fmt.Printf("your wish \"" + input + "\" will happen soon ")

}
