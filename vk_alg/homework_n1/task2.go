package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size, element int
	fmt.Scanln(&size)
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Scan(&element)
	fmt.Printf("%v", scientistsSolution(size, nums, element))

}

func scientistsSolution(size int, nums []int, elem int) string {
	res := ""
	numsCorrect := make([]int, 0, size)
	for i := 0; i < size; i++ {
		if nums[i] == elem {
			continue
		} else {
			numsCorrect = append(numsCorrect, nums[i])
		}
	}
	for i := 0; i < len(numsCorrect); i++ {
		if i != len(numsCorrect)-1 {
			res += strconv.Itoa(numsCorrect[i]) + " "
		} else {
			res += strconv.Itoa(numsCorrect[i])
		}
	}
	return res
}
