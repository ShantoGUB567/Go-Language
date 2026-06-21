package main

import "fmt"

func mapPractice() {
	student := map[string]int{
		"shanto": 26,
		"Robin":  25,
		"Shamim": 22,
	}
	fmt.Println(student) 

	// add new 
	student["Nafiz"] = 10 
	student["Himu"] = 15 

	delete(student, "Shamim") 

	for key, value := range student {
		fmt.Println(key, value)
	}
}