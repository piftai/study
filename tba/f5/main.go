package main

import (
	"fmt"
)

func main() {

	var n int
	var a, b int64
	fmt.Scan(&n, &a, &b)

	var s string
	fmt.Scan(&s)

	open, closed := 0, 0
	for _, c := range s {
		if c == '(' {
			open++
		} else {
			closed++
		}
	}

	cost := int64(0)
	if open > n {
		cost += int64(open-n) * b
		closed += open - n
		open = n
	} else if closed > n {
		cost += int64(closed-n) * b
		open += closed - n
		closed = n
	}

	balance := 0
	errors := 0
	for _, c := range s {
		if c == '(' {
			balance++
		} else {
			if balance > 0 {
				balance--
			} else {
				errors++
			}
		}
	}

	if a < b {
		cost += int64(errors/2) * a
	} else {
		cost += int64(errors) * b
	}

	fmt.Println(cost)
}
