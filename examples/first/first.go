package first

import (
	"fmt"
	"math"
)

func main() {
	welcome()
	fmt.Println("-------")

	// declaration typed variable
	var number1 float64 = 3.2
	var number2 float64 = 2
	// abbreviation of variable declaration
	result := sumFloat(number1, number2)
	// print with some args
	fmt.Println("Sum of two numbers: ", number1, " + ", number2, " = ", result)
	// Print Variable type
	fmt.Printf("Type result is %T\n", result)
	// compare with if's
	biggerThan(result, 3)

	fmt.Println("-------")
	// for example
	var factorial int = getFactorial(6)
	fmt.Println("(6)! = ", factorial)

	fmt.Println("-------")
	// for example
	resultSquare := math.Sqrt(25)
	fmt.Println("square root of 25 = ", resultSquare)

}

func welcome() {
	fmt.Println("Welcome to my first Code on Golang")
}

func sumFloat(x float64, y float64) float64 {
	return x + y
}

func biggerThan(x float64, limit float64) {
	if x > limit {
		fmt.Println("It's bigger than ", limit)
	} else {
		fmt.Println("It's not bigger than ", limit)
	}
}

func getFactorial(factor int) int {
	var result int = 1
	for i := factor; i > 0; i-- {
		result = result * i
	}
	return result
}
