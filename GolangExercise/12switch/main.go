package main

import (
	"fmt"
	"math/rand"
)

func main() {

	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("You can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and anouther dice throw")
	}
}
