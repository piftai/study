package main

import "fmt" // Print tests

func main() {
	fmt.Printf("TEST CASE 1: %v\n", isValid("()"))
	fmt.Printf("TEST CASE 2: %v\n", isValid("12"))
	fmt.Printf("TEST CASE 3: %v\n", isValid("(]"))
	fmt.Printf("TEST CASE 4: %v\n", isValid(" "))
	fmt.Printf("TEST CASE 5: %v\n", isValid("\n"))
	fmt.Printf("TEST CASE 6: %v\n", isValid(`0`))
	fmt.Printf("TEST CASE 7: %v\n", isValid(`127`))
	a := ""
	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Printf("Incorrect data in test case №%d", 8+i)
		} else {
			fmt.Printf("TEST CASE %d: %v\n", 8+i, isValid(a))
		}
	}
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]string, 0, len(s)/2) // created a slice with string values. cap of slice is length s / 2

	for i := 0; i < len(s)/2; i++ {
		stack = append(stack, string(s[i]))
	}

	for i := len(s) / 2; i < len(s); i++ {
		if len(stack) > 0 && (stack[len(stack)-1] == "(" && string(s[i]) == ")" || stack[len(stack)-1] == "[" && string(s[i]) == "]" || stack[len(stack)-1] == "{" && string(s[i]) == "}") {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true
}
