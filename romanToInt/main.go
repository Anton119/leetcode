package main

import "fmt"

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	arr := []string{}

	for i := 0; i < len(s); i++ {
		arr = append(arr, string(s[i]))
	}

	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && roman[arr[i]] < roman[arr[i+1]] {
			res -= roman[arr[i]]
		} else {
			res += roman[arr[i]]
		}
	}

	return res
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
