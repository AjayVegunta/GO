package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("welcome!")
	fmt.Println("enter the money you want to buy in")
	var total_money int
	fmt.Scanf("%v", &total_money)
	for total_money >= 0 {
		fmt.Println("Welcome to the game")
		fmt.Println("Enter the amount you want to bet")
		var bet int
		fmt.Scanf("%v", &bet)
		if total_money >= bet {
			fmt.Println("Select your bet \n1.red/black\n2.high or low\n3.number")
			var ch string
			fmt.Scanf("%s", &ch)
			randomInt := rand.Intn(36)
			total_money = total_money - bet
			if ch == "red" {
				if randomInt%2 == 0 {
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("you won 1:1 on what you have bet")
					total_money = total_money + (bet * 2)
					fmt.Println("your new balance is:", total_money)
				} else {
					fmt.Println("You Lost!")
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("Your current balance is", total_money)
				}

			} else if ch == "black" {
				if randomInt%2 != 0 {
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("you won 1:1 on what you have bet")
					total_money = total_money + (bet * 2)
					fmt.Println("your new balance is:", total_money)
				} else {
					fmt.Println("You Lost!")
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("Your current balance is", total_money)
				}
			} else if ch == "low" {
				if randomInt >= 1 && randomInt <= 18 {
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("you won 1:2 on what you have bet")
					total_money = total_money + bet + (bet * 2)
					fmt.Println("your new balance is:", total_money)
				} else {
					fmt.Println("You Lost!")
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("Your current balance is", total_money)
				}
			} else if ch == "high" {
				if randomInt >= 19 && randomInt <= 36 {
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("you won 1:2 on what you have bet")
					total_money = total_money + bet + (bet * 2)
					fmt.Println("your new balance is:", total_money)
				} else {
					fmt.Println("You Lost!")
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("Your current balance is", total_money)
				}
			} else if ch == "number" {
				fmt.Println("enter the specific number you want to bet")
				var specific int
				fmt.Scanf("%v", &specific)
				if specific == randomInt {
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("you won 1:35 on what you have bet")
					total_money = total_money + bet + (bet * 35)
					fmt.Println("your new balance is:", total_money)

				} else {
					fmt.Println("You lost!")
					fmt.Println("the number on the wheel is", randomInt)
					fmt.Println("Your current balance is", total_money)
				}

			}
		} else {
			fmt.Println("You dont have that much money")
			fmt.Println("Your current balance is", total_money)
			fmt.Println("Thank You! Visit again")
			break
		}
	}
}
