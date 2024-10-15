package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 0, 1, 0, 3, 12}
	fmt.Printf("%s", solution(6, nums))
}

func solution(n int, nums []int) string {
	zeroAt := 0
	for current, value := range nums {
		if value != 0 {
			nums[zeroAt], nums[current] = nums[current], nums[zeroAt]
			zeroAt++
		}
	}
	word := ""
	for i, v := range nums {
		if i != len(nums)-1 {
			word += strconv.Itoa(v) + " "
		} else {
			word += strconv.Itoa(v)
		}
	}
	return word
}
