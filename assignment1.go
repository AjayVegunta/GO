package main

import "fmt"

// Factorial Function
func factorial(num int) int {
	if num == 0 || num == 1 {
		return num
	} else {
		return num * factorial(num-1)
	}

}
func main() {

	//Arthmetic Operators
	a := 3
	b := 4
	fmt.Println("Sum is", a+b)
	fmt.Println("Difference is ", a-b)
	fmt.Println("Product is", a*b)
	fmt.Println("Division is", a/b)
	fmt.Println("remainder is", a%b)

	greet := "hello"
	name := "Ajay"

	if greet == "hello" {
		fmt.Println("Welcome", name)
	} else {
		fmt.Println("Invalid")
	}
	//Relational Operators
	c := 10
	d := 15
	if c == d {
		fmt.Println("both equal")
	} else if c > d {
		fmt.Println("c is bigger")
	} else {
		fmt.Println("d is bigger")
	}
	//Factorial of a number by calling factorial function
	fmt.Println("the factorial of 5 is", factorial(5))

	//Tried to demonstrate switch and logical operators
	var x bool = true
	var y bool = true
	switch true {
	case x == true && y == true:
		fmt.Println("both true case")
	case x != true || y != true:
		fmt.Println("both false case")
	default:
		fmt.Println("Invalid")
	}
	//Printing an array using for loop
	g := []int{20, 30, 40, 50}
	for j := 0; j < 4; j++ {
		fmt.Println(g[j])
	}

}
