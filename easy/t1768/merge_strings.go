package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	ans := strings.Builder{}
	l := 0
	r := 0
	for l < len(word1) && r < len(word2) {
		ans.WriteByte(word1[l])
		ans.WriteByte(word2[r])
		l++
		r++
	}
	return ans.String() + word1[l:] + word2[r:]
}
