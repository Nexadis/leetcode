package main

import "strings"

func reverseParentheses(s string) string {
	sb := []byte(s)
	for i := 0; i < len(sb); i++ {
		r := sb[i]
		switch r {
		case '(':
			reversed := reverseParentheses(string(sb[i+1:]))
			for j := range reversed {
				sb[i+j+1] = reversed[j]
			}
			i += len(reversed) + 1
		case ')':
			return string(reverse(sb[:i]))
		}
	}
	s = string(sb)
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	return s
}

func reverse(sb []byte) []byte {
	for i := 0; i < len(sb)/2; i++ {
		sb[i], sb[len(sb)-i-1] = sb[len(sb)-i-1], sb[i]
	}
	return sb
}
