package main

import "fmt"

func factorial(num int) int {
	if num == 0 || num == 1 {
		return num
	} else {
		return num * factorial(num-1)
	}
}
func main() {
	fmt.Println("the factorial of 5 is ", factorial(5))
	fmt.Println("the facroial of 10 is ", factorial(10))
}
