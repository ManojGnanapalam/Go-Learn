package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter you age ! ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	newAge, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		if newAge >= 18 {
			fmt.Printf("You are above " + input + " ,you vote")
		}
	}
}
