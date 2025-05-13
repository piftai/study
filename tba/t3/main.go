package main

import "fmt"

// tests
// 21
// 0 125 250 442 2 2 1442 442 250 2 2 2 2 2 1 1 1 3 3 3 3

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	array := make([]int, n)
	for i := 0; i < len(array); i++ {
		if _, err := fmt.Scan(&array[i]); err != nil {
			return
		}
	}
	uniq := make(map[int]struct{})
	for i := 0; i < n; i++ {
		current := array[i]
		for {
			if _, exist := uniq[current]; !exist {
				uniq[current] = struct{}{}
				break
			} else {
				current = current / 2
			}
			if current == 0 {
				uniq[current] = struct{}{}
				break
			}
		}
	}
	fmt.Println(len(uniq), uniq)

}
