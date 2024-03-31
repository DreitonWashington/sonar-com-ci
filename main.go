package main

import "fmt"

func main() {
	fmt.Print(soma(2, 2))
	fmt.Print(sub(2, 2))
	fmt.Print(multi(2, 2))
	fmt.Print(divi(2, 2))
}

func soma(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func divi(a, b int) int {
	return a / b
}
