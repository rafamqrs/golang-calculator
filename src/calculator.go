package main

import "fmt"

func main() {
	fmt.Println("Please type the first number")
	var number1 int
	fmt.Scan(&number1)
	
	fmt.Println("Please type the second number")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("the sum is", number1 + number2)

}

