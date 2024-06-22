package main

import "fmt"

func main() {
	nameList := make(map[int]string)

	nameList[1] = "naveen"
	nameList[2] = "vaiper"
	nameList[3] = "keerthi"
	nameList[4] = "mag"

	fmt.Println(nameList)
	fmt.Println(nameList[1])

	delete(nameList, 4)
	fmt.Println(nameList)

	for _, value := range nameList {
		fmt.Println(value)
	}
}
