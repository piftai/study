package main

import "fmt"

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
	cnt := 0
	for i := 0; i < len(array)-2; i++ {
		a := array[i]
		b := array[i+1]
		c := array[i+2]
		if b-a == c-b {
			cnt += len(array) - i - 2
		}
	}
	fmt.Println(cnt)
}
