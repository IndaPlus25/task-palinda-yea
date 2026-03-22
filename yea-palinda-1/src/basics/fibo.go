package main

import "fmt"

func fibonacci() func() int {
	previous, number := 0, 1
	return func() int {
		current := previous
		previous, number = number, number+previous
		return current
	}
}

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
