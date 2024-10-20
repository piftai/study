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
	sort.Ints(candies)
	maxCandy := candies[len(candies)-1]
	fmt.Printf("maxCandy: %d\n", maxCandy)
	res := make([]bool, 0, len(candies)+1)
	for _, v := range candies {
		if v+extraCandies >= maxCandy {
			fmt.Printf("%v because v+extraCandies =")
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
