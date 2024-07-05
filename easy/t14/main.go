package main

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		s := strs[i]
		l := min(len(prefix), len(s))
		j := 0

		for j < l && prefix[j] == s[j] {
			j++
		}

		prefix = prefix[:j]
	}
	return prefix
}
