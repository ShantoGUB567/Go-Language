package main

import (
	"fmt"
)

func array() {
	// var numbers [5]int // blank array create
	numbers := [5]int{20, 42, 83, 97, 17}
	// fmt.Println(numbers[4])
	fmt.Println(numbers[4])
	fmt.Println(numbers[len(numbers)-1])
	numbers[2] = 33 

	for number := 0; number < len(numbers); number++ {
		fmt.Println(numbers[number])
	}
}
