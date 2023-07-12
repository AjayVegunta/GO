package main

import "fmt"

type ajay struct {
	x int
	y int
}

func add(a ajay) int {
	return a.x + a.y
}
func sub(a ajay) int {
	return a.x - a.y
}
func mul(a ajay) int {
	return a.x * a.y
}
func div(a ajay) int {
	return a.x / a.y
}
func rem(a ajay) int {
	return a.x % a.y
}
func logical(num1, numm2 int) string {
	var k *int
	var l *int
	k = &num1
	l = &numm2
	if *k > *l && *k%2 == 0 {
		return "num1 is greater than num2 and is a even number"
	} else if *k == *l {
		return "Both are equal"
	} else if *k < *l && *l%2 == 0 {
		return "num2 is graeter than num1 and is a even number"
	} else {
		return "invalid"
	}

}
func main() {
	a := ajay{10, 20}
	fmt.Println("Enter the Operation you want to perform\n1.addition\n2.subtraction\n3.multiplication\n4.division\n5.remainder")
	var operation string
	fmt.Scanln(&operation)
	switch operation {
	case "addition":
		fmt.Println(add(a))
	case "subtraction":
		fmt.Println(sub(a))
	case "multiplication":
		fmt.Println(mul(a))
	case "division":
		fmt.Println(div(a))
	case "remainder":
		fmt.Println(rem(a))
	default:
		fmt.Println("Invalid choice")
	}
	fmt.Println(logical(34, 54))
}
