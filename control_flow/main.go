package main

import "fmt"

func main() {
	// displayEvenNum()
	// displaySum()
	// multiplicationTable()

	greet("Shanto")
	var add, mul int
	var div float64
	add, mul, div = calculation(46, 21)
	fmt.Println("Add: ",add, " Mul: ",mul, " Div: ", div)
	fmt.Println("Factorial of 5: ", factorial(5))
}