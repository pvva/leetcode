package main

func main() {
	println(maxSubstr("abbca"))
	println(maxSubstr("abba"))
	println(maxSubstr("abcdcbad"))
}

/*
Given a string find maximum substring that contains unique characters.

Input: abbca
Output: bca

Input: abba
Output: ab or ba
*/
func maxSubstr(s string) string {
	start := 0
	end := 0
	ts := 0
	i := 0
	l := len(s)
	set := make(map[byte]bool)

	for i = 0; i < l; i++ {
		if set[s[i]] {
			set = make(map[byte]bool)
			start = ts
			end = i
			ts = i
		}
		set[s[i]] = true
	}
	if i-ts > end-start {
		start = ts
		end = i
	}

	return s[start:end]
}
