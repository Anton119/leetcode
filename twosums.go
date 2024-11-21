package main

import "fmt"

func twoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	arr := []int{}
	for i, number := range nums {
		_, ok := m[number]
		if ok {
			return []int{m[number], i}
		}
		m[target-number] = i
	}
	return arr
	//com
}

func main() {
	nums := []int{2, 7, 5, 11, 4}
	target := 6
	fmt.Println(twoSumHash(nums, target))
}
