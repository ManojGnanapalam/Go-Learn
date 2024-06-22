package main

import (
	"fmt"
	"sort"
)

func main() {
	var nameList = []string{"naveen", "keerthi", "vaiper"}
	fmt.Println(nameList)

	nameList = append(nameList, "spark")

	fmt.Println(nameList)

	fmt.Printf("type of the name list :%T\n", nameList)

	var highscores = make([]int, 4)

	highscores[0] = 123
	highscores[1] = 233
	highscores[2] = 332
	highscores[3] = 135
	//highscores[4] = 1366
	fmt.Println(highscores)
	sort.Ints(highscores)
	fmt.Println(highscores)

	//  how to remove a value from slices based on index

	index := 2
	highscores = append(highscores[:index], highscores[index+1:]...)

	fmt.Println(highscores)
	for value := range nameList {
		fmt.Println(nameList[value])
	}

}
