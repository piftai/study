package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cnt int
	fmt.Scan(&cnt)
	offers := make(map[string][]string, 3)
	for i := 0; i < cnt; i++ {
		for j := 1; j < 4; j++ {
			offers[strconv.Itoa(j)] = append(offers[strconv.Itoa(j)], str)
		}
	}
}
