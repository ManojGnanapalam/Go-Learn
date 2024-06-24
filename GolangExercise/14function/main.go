package main

import "fmt"

func main() {
	res, ans, supAns := adder(2, 3)
	fmt.Println(res, ans, supAns, superAdder(2, 3, 4))

	superRes := superAdder(2, 3, 4, 5, 2)
	fmt.Println(superRes)
}

func adder(a int, b int) (int, string, string) {
	return a + b, "ans for adder", "super funciton "
}

func superAdder(values ...int) int {
	ans := 0

	for _, num := range values {
		ans += num
	}

	return ans
}
