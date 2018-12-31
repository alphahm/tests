// Fibonacci Recursion test

package main

import "fmt"

func main() {
	fmt.Println(fibonacci(40))
}

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
