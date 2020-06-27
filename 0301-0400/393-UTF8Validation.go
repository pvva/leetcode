package main

func main() {
	println(validUtf8([]int{197, 130, 1}), true)
	println(validUtf8([]int{235, 140, 1}), false)
	println(validUtf8([]int{255}), false)
	println(validUtf8([]int{230, 136, 145}), true)
	println(validUtf8([]int{145}), false)
	println(validUtf8([]int{240, 162, 138, 147}), true)
	println(validUtf8([]int{194, 155, 231, 184, 185, 246, 176, 131, 161, 222, 174, 227, 162, 134, 241, 154, 168, 185,
		218, 178, 229, 187, 139, 246, 178, 187, 139, 204, 146, 225, 148, 179, 245, 139, 172, 134, 193, 156, 233, 131,
		154, 240, 166, 188, 190, 216, 150, 230, 145, 144, 240, 167, 140, 163, 221, 190, 238, 168, 139, 241, 154, 159,
		164, 199, 170, 224, 173, 140, 244, 182, 143, 134, 206, 181, 227, 172, 141, 241, 146, 159, 170, 202, 134, 230,
		142, 163, 244, 172, 140, 191}), true)
}

/*
https://leetcode.com/problems/utf-8-validation/

A character in UTF8 can be from 1 to 4 bytes long, subjected to the following rules:

For 1-byte character, the first bit is a 0, followed by its unicode code.
For n-bytes character, the first n-bits are all one's, the n+1 bit is 0, followed by n-1 bytes with
most significant 2 bits being 10.
This is how the UTF-8 encoding would work:

   Char. number range  |        UTF-8 octet sequence
      (hexadecimal)    |              (binary)
   --------------------+---------------------------------------------
   0000 0000-0000 007F | 0xxxxxxx
   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
Given an array of integers representing the data, return whether it is a valid utf-8 encoding.

Note:
The input is an array of integers. Only the least significant 8 bits of each integer is used to store the data.
This means each integer represents only 1 byte of data.

Example 1:

data = [197, 130, 1], which represents the octet sequence: 11000101 10000010 00000001.

Return true.
It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte character.
Example 2:

data = [235, 140, 4], which represented the octet sequence: 11101011 10001100 00000100.

Return false.
The first 3 bits are all one's and the 4th bit is 0 means it is a 3-bytes character.
The next byte is a continuation byte which starts with 10 and that's correct.
But the second continuation byte does not start with 10, so it is invalid.
*/

func validUtf8(data []int) bool {
	for idx := 0; idx < len(data); {
		v := data[idx]
		if v&248 == 248 {
			return false
		}
		n := 1
		switch {
		case v&240 == 240: // 11110000
			n = 4
		case v&224 == 224: // 11100000
			n = 3
		case v&192 == 192: // 11000000
			n = 2
		case v&128 == 128: // 10000000
			return false
		}

		if idx+n > len(data) {
			return false
		}

		for i := 1; i < n; i++ {
			if data[idx+i]&64 != 0 || data[idx+i]&128 != 128 {
				return false
			}
		}
		idx += n
	}

	return true
}
