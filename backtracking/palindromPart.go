package main

import "fmt"

func main() {
	//fmt.Println(partition("aab"))
	fmt.Println(partition("aaaa"))
	//fmt.Println(partition("seeslaveidemonstrateyetartsnomedievalsees"))

	//fmt.Println(allSubsets("132"))
}

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
*/
func partition(s string) [][]string {
	res := [][]string{}
	partitionBT(&res, nil, make(map[string]bool), s, 0)

	return res
}

func partitionBT(res *[][]string, tlist *[]string, palindromes map[string]bool, s string, start int) {
	if tlist == nil {
		tlist = &[]string{}
	}
	if start == len(s) {
		clist := make([]string, len(*tlist))
		copy(clist, *tlist)
		*res = append(*res, clist)
	} else {
		for i := start; i < len(s); i++ {
			s0 := s[start : i+1]
			if isPalindrome(s0, palindromes) {
				*tlist = append(*tlist, s0)
				partitionBT(res, tlist, palindromes, s, i+1)
				t := *tlist
				*tlist = t[:len(t)-1]
			}
		}
	}
}

func partitionBF(s string) [][]string {
	return parts(s, make(map[string]bool))
}

func parts(s string, palindromes map[string]bool) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res := [][]string{}
	if isPalindrome(s, palindromes) {
		res = append(res, []string{s})
	}

	for k := 1; k < len(s); k++ {
		s0 := s[0:k]
		if !isPalindrome(s0, palindromes) {
			continue
		}

		rParts := parts(s[k:], palindromes)

		for i := 0; i < len(rParts); i++ {
			res = append(res, append([]string{s0}, rParts[i]...))
		}
	}

	return res
}

func isPalindrome(s string, palindromes map[string]bool) bool {
	isp, ex := palindromes[s]
	if !ex {
		isp = true
		l := len(s)
		for i := 0; i < l>>1; i++ {
			if s[i:i+1] != s[l-i-1:l-i] {
				isp = false
				break
			}
		}
		palindromes[s] = isp
	}

	return isp
}

func allSubsets(s string) [][]string {
	res := [][]string{}
	bt(&res, nil, s, 0)

	return res
}

func bt(res *[][]string, tlist *[]string, s string, start int) {
	if tlist != nil {
		clist := make([]string, len(*tlist))
		copy(clist, *tlist)
		*res = append(*res, clist)
	} else {
		tlist = &[]string{}
		*res = append(*res, []string{})
	}

	for i := start; i < len(s); i++ {
		*tlist = append(*tlist, s[i:i+1])
		bt(res, tlist, s, i+1)
		t := *tlist
		*tlist = t[:len(t)-1]
	}
}
