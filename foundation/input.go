package main

import "fmt"

func inputFromConsole() {
	var name string
	var mass float64
	var weight float64

	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Print("Enter your weight : ")
	fmt.Scanln(&weight)

	const gravity float64 = 9.8
	mass = weight / gravity 
	
	fmt.Println("Name: ", name, "\nWeight: ", weight, "\nMass: ", mass)

	var num1, num2 int 
	fmt.Print("Enter 2 integer number: ")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Printf("Addition: %d\n", num1 + num2)
	fmt.Printf("Subtraction: %d\n", num1 - num2)
	fmt.Printf("Multiplecation: %d\n", num1 * num2)
	fmt.Printf("Division: %f\n", num1 / num2)
}