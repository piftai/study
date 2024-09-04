package main

import "fmt"

func main() {
	str := []string{"a", "a", "a", "aabam"}
	str1 := []string{"floor", "flower", "flight"}
	str2 := []string{"ab", "ab", "ab"}
	fmt.Printf("TEST CASE 1: %s\n", longestCommonPrefix(str))  // correct answer "a"
	fmt.Printf("TEST CASE 2: %s\n", longestCommonPrefix(str1)) // correct answer "fl"
	fmt.Printf("TEST CASE 3: %s\n", longestCommonPrefix(str2)) // correct answer "ab"

}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i := range len(strs[0]) {
		for _, v := range strs {
			if i == len(v) || v[i] != strs[0][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}
	return prefix
}
