package main

import "fmt"

func check_even(num int) string {
	if num%2 == 0 {
		return "even number"
	} else {
		return "odd number"
	}

}
func main() {
	a := check_even(6)
	fmt.Println(a)
}
