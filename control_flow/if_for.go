package main

import "fmt"

// Even numbers 1 to 50
func displayEvenNum(){
	fmt.Print("Even Numbers (1 to 50) : ") 

	for i :=1 ; i <=50; i++ {
		if i%2 == 0 {
			fmt.Print(i , " ")
		}
	}
	fmt.Println()
}

// Sum of 1 to 100
func displaySum() {
	sum := 0

	for i :=1 ; i <=100; i++ {
		// fmt.Print(i, " ")
		sum += i
	}
	fmt.Println("Sum (1 to 100) : ", sum)
}

func multiplicationTable() {
	num :=5 
	fmt.Println("Multiplication Table of ", num)

	for i := 1; i <= 10; i++ {
		fmt.Println(num, " * ", i, " = ", num*i)
	}
}