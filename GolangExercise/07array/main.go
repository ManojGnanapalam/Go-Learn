package main

import "fmt"

func main() {
	var nameList = [4]string{"naveen", "vaigunth", "keerthi"}

	fmt.Println("users :", nameList)
	fmt.Println("Length of nameList :", len(nameList))
	fmt.Printf("type of the name list :%T\n", nameList)
}
