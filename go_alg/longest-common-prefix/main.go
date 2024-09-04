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
	prefix := strs[0]
	var cnt int
	if len(prefix) == 1 {
		for _, val := range strs {
			if string(val[0]) == prefix {
				cnt++
			}
		}
	}
	if cnt == len(strs) {
		return prefix
	}
	for i := len(prefix) - 1; i >= 0; i-- { // делаем массив, который будет итерироваться по срезу prefix
		prefix = prefix[:i]
		fmt.Printf("i: %d, prefix: %s\n", i, prefix)
		cnt = 0
		for _, v := range strs { // массив, который сверяет срез каждого элемента strs со срезом prefix, равны ли они друг другу
			if len(v) >= i && v[:i] == prefix { // чтобы не было out of range - мы должны
				// проверять есть ли в элементе столько же или больше символов, чем в префиксе
				cnt++
			} else {
				break // отсекаем лишнее. если понимаем, что префикс уже не подходит
			}
		}
		if cnt == len(strs) {
			return prefix
		}
	}
	return prefix
}
