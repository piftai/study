package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size int
	fmt.Scanln(&size)
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Printf("%s", solution(size, nums))
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
