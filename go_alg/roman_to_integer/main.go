package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	m := make(map[string]int, 7)
	m["I"] = 1 // 1000 I = 0
	m["V"] = 5 // I = 1
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	var cnt int
	for i, v := range s {

		if i != 0 {
			if m[string(s[i-1])] >= m[string(s[i])] { // 0 >= 1 XIV 11
				cnt += m[string(v)]
			} else {
				cnt += m[string(v)]
				cnt -= 2 * m[string(s[i-1])]
			}
		} else {
			cnt += m[string(v)]
		}
	}
	return cnt
}
