package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func calculation(a, b int) (int, int, float64) {
	return a+b, a*b, float64(a/b)
}

func factorial(n int) int {
	if n == 1 {
		return 1 
	} 
	return n * factorial(n-1)
}