package main

type chars [26]int

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sc := chars{}
	st := chars{}
	for _, c := range s {
		sc[c-'a']++
	}
	for _, c := range t {
		st[c-'a']++
	}
	return sc == st
}
