package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	exWord := strs[0]

	for j := 0; len(exWord) != 0; j-- {
		var cnt int
		var curr string
		for i, v := range strs {
			curr = exWord[:len(v)]
			if exWord[:len(v)] == v {
				cnt++
			} else {
				break
			}
		}

	}
}
