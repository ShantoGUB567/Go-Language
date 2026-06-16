package main

import "fmt"

func variableDisplay() {
	var name , university string = "Shanto" , "GUB"
	var age int = 25
	var cgpa float32 = 3.44
	fmt.Println(name, age, university, cgpa)

	var country = "Bangladesh" 
	fmt.Println(country)

	city := "dhaka"
	fmt.Println(city)

	x, y, z := 5,6,7 
	println(x,y,z)

	var salary int 
	var company string 
	var designation string 
	fmt.Println(salary, company, designation)
	fmt.Printf("%T\n", city)
}