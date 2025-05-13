package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	for i := 0; i < 4; i++ {
		newstr := str[:i] + str[i+1:]
		isPal := true
		left, right := 0, len(newstr)-1
		for left != right {
			if newstr[left] != newstr[right] {
				isPal = false
				break
			} else {
				left++
				right--
			}
		}
		if isPal {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
