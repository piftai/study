package main

import "fmt"

func main() {
	fmt.Println()
	var cnt int
	fmt.Scan(&cnt)
	var str string
	answers := make([]string, 0, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&str)
		answers = append(answers, canBeConstructed(str))
	}
	for _, v := range answers {
		fmt.Println(v)
	}
}

func canBeConstructed(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return "YES"
	}
	if str[0] != str[len(str)-1] {
		return "NO"
	}
	src := str[0]

	for i := 1; i < len(str); i++ {
		if str[i] != src {
			if i+1 < len(str) && str[i+1] == src {
				continue
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}
