package main

import "fmt"

func main() {
	var length, element int
	_, err := fmt.Scan(&length)
	if err != nil {
		return
	}
	nums := make([]int, 5)
	for i := 0; i < length; i++ {
		_, err = fmt.Scan(&nums[i])
		if err != nil {
			return
		}
	}
	_, err = fmt.Scan(&element)
	if err != nil {
		return
	}
	fmt.Printf("%d", solution(length, nums, element))
}

func solution(len int, nums []int, element int) int {
	var cnt int = 0
	for _, v := range nums {
		if v != element {
			cnt++
		}
	}
	return cnt
}
