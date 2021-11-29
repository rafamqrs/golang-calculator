package main

import "fmt"

func main() {
	fmt.Println("Please type the first number")
	var number1 int
	fmt.Scan(&number1)
	
	fmt.Println("Please type the second number")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("the sum is", sum(number1, number2))
	
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
