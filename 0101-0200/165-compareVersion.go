package main

func main() {
	println(compareVersion("0.1", "1.1"), -1)
	println(compareVersion("1.0.1", "1"), 1)
	println(compareVersion("7.5.2.4", "7.5.3"), -1)
	println(compareVersion("1.01", "1.001"), 0)
}

/*
https://leetcode.com/problems/compare-version-numbers/

Compare two version numbers version1 and version2.
If version1 > version2 return 1; if version1 < version2 return -1;otherwise return 0.

You may assume that the version strings are non-empty and contain only digits and the . character.

The . character does not represent a decimal point and is used to separate number sequences.

For instance, 2.5 is not "two and a half" or "half way to version three", it is the fifth second-level revision
of the second first-level revision.

You may assume the default revision number for each level of a version number to be 0. For example, version number
3.4 has a revision number of 3 and 4 for its first and second level revision number. Its third and fourth level
revision number are both 0.
*/

func compareVersion(version1 string, version2 string) int {
	idx1 := 0
	idx2 := 0

	for idx1 < len(version1) || idx2 < len(version2) {
		v1 := 0
		v2 := 0
		v1, idx1 = getNextVersionChunk(version1, idx1)
		v2, idx2 = getNextVersionChunk(version2, idx2)

		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	return 0
}

func getNextVersionChunk(s string, idx int) (int, int) {
	value := 0

	for idx < len(s) {
		if s[idx] == '.' {
			idx++
			break
		}
		value = value*10 + int(s[idx]-'0')
		idx++
	}

	return value, idx
}
