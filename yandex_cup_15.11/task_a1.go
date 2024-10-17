package main

import (пш
	"fmt"
	"math"
)

func main() {
	var a, b, c, n uint64
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}
	list := make([]uint64, 0, n)
	for i := uint64(1); float64(i) < math.Pow(10, 18); i++ {
		if i%a == 0 && i%b == 0 && i%c != 0 || i%a == 0 && i%c == 0 && i%b != 0 || i%b == 0 && i%c == 0 && i%a != 0 {
			list = append(list, i)
			if uint64(len(list)) > n && uint64(len(list)) < n+2 {
				fmt.Printf("%d", list[n-1])
				return
			}
		}
	}
	fmt.Print("-1")
}
