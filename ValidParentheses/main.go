package main

import "fmt"

func main() {
	str := "]"
	fmt.Println(str, isValid(str))
	str = "()[]{}"
	fmt.Println(str, isValid(str))
	str = "(]"
	fmt.Println(str, isValid(str))
	str = "([])"
	fmt.Println(str, isValid(str))
	str = "("
	fmt.Println(str, isValid(str))
}

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	count := len(s)

	for i := 0; i < count; i++ {
		ch := s[i]
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			tidx := len(stack) - 1
			top := stack[tidx]
			if top != m[ch] {
				return false
			}
			stack = stack[:tidx]
		}
	}
	return len(stack) == 0

}
