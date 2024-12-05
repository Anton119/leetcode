package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	sort.Strings(strs)

	minLength := len(strs[0])
	for _, word := range strs {
		if minLength > len(word) {
			minLength = len(word)
		}
	}

	for i := 0; i < minLength; i++ {
		if strs[0][i] != strs[len(strs)-1][i] {
			return res
		}
		res += string(strs[0][i])
	}

	return res
}

func main() {
	strs := []string{"fliwin", "flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
