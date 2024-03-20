package main

import "fmt"

func main() {
	fmt.Println("Hello, Worlds!")
	fmt.Println(Calculate(4))
}

func Calculate(x int) (result int) {
	result += x
	return result
}
