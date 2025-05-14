package main

import "fmt"

func main() {
	fmt.Printf("Test foo with 1:%d\n", foo(1))
	fmt.Printf("Test foo with 2:%d\n", foo(2))
	fmt.Printf("Test foo with 3:%d\n", foo(3))
	fmt.Printf("Test foo with 4:%d\n", foo(4))
	fmt.Printf("Test foo with 5:%d\n", foo(5))
}

func foo(n int) int {
	if n == 1 {
		return 1
	}
	return n * foo(n-1)
}
