package main

import "fmt"

func greet() {
	fmt.Println("Hi, your required sum is:")
}

func sum(a int, b int) int {
	return a + b
}

func main() {
	greet()
	fmt.Print(sum(5, 6))
}
