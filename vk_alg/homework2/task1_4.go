package main

import "fmt"

func main() {
	var size int
	fmt.Scanln(&size)
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Printf("%v", maxNumber(nums))
}

func maxNumber(nums []int) int {
	var maximum int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	return maximum
}
