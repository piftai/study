package main

import (
	"fmt"
	"sort"
)

func main() {
	var array = []int{2, 3, 5, 1, 3}

	var exCandies int = 3
	fmt.Printf("TEST 1: %v\n", kidsWithCandies(array, exCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	candiesSorted := make([]int, len(candies))
	copy(candiesSorted, candies)
	sort.Ints(candiesSorted)
	maxCandy := candiesSorted[len(candiesSorted)-1]
	fmt.Printf("maxCandy: %d\n", maxCandy)
	res := make([]bool, 0, len(candies)+1)
	for _, v := range candies {
		if v+extraCandies >= maxCandy {
			fmt.Printf("%v because v = %v extraCandies = %v\n", true, v, extraCandies)
			res = append(res, true)
		} else {
			fmt.Printf("%v because v = %v extraCandies = %v\n", false, v, extraCandies)
			res = append(res, false)
		}
	}
	return res
}
