package main

import "fmt" 

func main() {
	welcome()
	// declaration typed variable
	var number1 float64 = 3.2
	var number2 float64 = 2
	// abbreviation of variable declaration
	result := sumFloat(number1, number2)
	// print with some args
	fmt.Println("Sum of two numbers: ",number1," + ",number2," = ", result)
	// Print Variable type
	fmt.Printf("Result is %T\n", result)
}

func welcome() {
	fmt.Println("Welcome to my first Code on Golang:")
}

func sumFloat(x float64, y float64) float64 {
	return x + y
}
