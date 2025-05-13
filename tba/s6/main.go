package main

import (
	"fmt"
	"math"
)

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
	for len(array) > 1 {
		maxDiff := 0
		pos := 0

		for i := 0; i < len(array)-1; i++ {
			diff := int(math.Abs(float64(array[i] - array[i+1])))
			if diff > maxDiff {
				maxDiff = diff
				pos = i
			}
		}

		cnt += maxDiff

		array = append(array[:pos], array[pos+2:]...)
	}

	fmt.Println(cnt)
}
