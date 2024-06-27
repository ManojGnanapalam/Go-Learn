package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	checkError(err)

	fmt.Println(string(databyte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
