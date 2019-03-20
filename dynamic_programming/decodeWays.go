package main

func main() {
	println(numDecodings("226"), 3)
	println(numDecodings("12"), 2)
	println(numDecodings("1226"), 5)
	println(numDecodings("1242657"), 6)
	println(numDecodings("1272657"), 4)
	println(numDecodings("1"), 1)
	println(numDecodings("0"), 0)
	println(numDecodings("10"), 1)
	println(numDecodings("101"), 1)
	println(numDecodings("110"), 1)
	println(numDecodings("1010"), 1)
}

func numDecodings(s string) int {
	if s[0] == 48 {
		return 0
	}

	prev := 1
	curr := 1

	for i := 0; i < len(s); i++ {
		t := 0
		// if current digit > 0 - add current to new amount
		if s[i] > 48 {
			t = curr
		}
		if i > 0 {
			v := int(s[i-1])*10 - 480 + int(s[i]) - 48
			// if with previous digit we have 10..26 - add previous amount
			if v > 9 && v <= 26 {
				t += prev
			}
		}
		prev, curr = curr, t
	}

	return curr
}
