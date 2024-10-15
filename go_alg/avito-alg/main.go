package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4}
	var b = []int{4, 4, 4}
	fmt.Println(sumArr(a, b))
}

func sumArr(arr1 []int, arr2 []int) []int {
	ret := make([]int, len(arr1))
	curr := 0
	if len(arr1) > len(arr2) {
		j := len(arr2) - 1
		for i := len(arr1) - 1; i >= 0; i-- {
			if j >= 0 {
				if arr1[i]+arr2[j]+curr >= 10 {
					ret = append(ret, arr1[i]+arr2[j]+curr%10)
					curr += 1
				} else {
					ret = append(ret, arr1[i]+arr2[j]+curr)
					curr = 0
				}
				j--
			}
		}
	}
	return ret
}
