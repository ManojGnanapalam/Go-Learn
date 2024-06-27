package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This message need to write on text file"

	file, err := os.Create("./newFile.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println("length is ", length)
	defer file.Close()

	databyte, err := os.ReadFile("./newFile.txt")
	checkError(err)
	fmt.Println(string(databyte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
