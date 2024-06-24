package main

import "fmt"

func main() {
	var nameList = []string{"vaiper", "naveen", "keethi", "spark"}

	for i := range nameList {
		fmt.Println(nameList[i])
	}

	for _, name := range nameList {
		fmt.Println(name)
	}

	randValue := 1

	for randValue < 10 {
		if randValue == 3 {
			goto res
		}
		if randValue == 5 {
			randValue++
			continue
		}
		fmt.Println("value is ", randValue)
		randValue++
	}

res:
	fmt.Println("label is here")
}
