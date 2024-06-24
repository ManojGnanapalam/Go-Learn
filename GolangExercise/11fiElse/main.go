package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the nunber")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	trimInput := strings.TrimSpace(input)

	if number, _ := strconv.Atoi(trimInput); number == 0 {
		fmt.Printf("%v is zero\n", number)
	} else if number%2 == 0 {
		fmt.Printf("%v is even\n", number)
	} else {
		fmt.Printf("%v is odd\n", number)
	}
}
