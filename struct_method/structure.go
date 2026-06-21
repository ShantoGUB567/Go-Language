package main

import "fmt"

func structure() {
	type employee struct {
		name   string
		age    int
		salary int
	}

	e1 := employee{
		name:   "Shanto",
		age:    25,
		salary: 50000,
	}

	fmt.Println(e1)
	e1.age = 26 
	fmt.Println(e1.age)

	var e2 employee 
	e2.name = "Robin"
	e2.age= 24 
	e2.salary = 62000 

	fmt.Println(e1.salary)
	fmt.Println(e2.name) 

	employees := []employee{
		{name: "Sanjid", age: 25, salary: 40000},
		{name: "Antu", age: 26, salary: 52000},
	}

	for _, emp := range employees {
		fmt.Println(emp.name, emp.age, emp.salary)
	}
}