package main

import "fmt"

func main() {
	var a map[string]int = map[string]int{
		"ajay":  18,
		"vijay": 20,
		"rahul": 33,
	}
	for name, age := range a {
		fmt.Println("name is", name, "age is", age)
	}
}
