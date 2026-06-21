package main

import "fmt"

func slice() {
	st_id := []int{3, 7, 9}
	fmt.Println(st_id)

	// Append 
	st_id = append(st_id, 4)
	st_id = append(st_id, 9)
	fmt.Println(st_id)

	fmt.Println("Lenth of  Slice: ", len(st_id))
	fmt.Println("Capacity of  Slice: ", cap(st_id))

	// Range Loop 
	for index, value := range st_id {
		fmt.Println(index, value)
	}

	numbers := [5]int{20, 42, 83, 97, 17} 
	number := numbers[1:4]
	fmt.Println(numbers)
	fmt.Println(number)
}