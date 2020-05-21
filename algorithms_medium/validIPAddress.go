package main

func main() {
	println(validIPAddress("172.16.254.1"))
	println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))

	println(validIPAddress("256.256.256.256"))
	println(validIPAddress("01.01.01.01"))
	println(validIPAddress("02001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	println(validIPAddress("2001:0db8:85a3:0:0000:8a2e:0370:7334"))
	println(validIPAddress("2001:0db8:85a3:0::8a2e:0370:7334"))
}

/*
https://leetcode.com/problems/validate-ip-address/

Write a function to check whether an input string is a valid IPv4 address or IPv6 address or neither.

IPv4 addresses are canonically represented in dot-decimal notation, which consists of four decimal numbers,
each ranging from 0 to 255, separated by dots ("."), e.g.,172.16.254.1;

Besides, leading zeros in the IPv4 is invalid. For example, the address 172.16.254.01 is invalid.

IPv6 addresses are represented as eight groups of four hexadecimal digits, each group representing 16 bits.
The groups are separated by colons (":"). For example, the address 2001:0db8:85a3:0000:0000:8a2e:0370:7334
is a valid one. Also, we could omit some leading zeros among four hexadecimal digits and some low-case characters
in the address to upper-case ones, so 2001:db8:85a3:0:0:8A2E:0370:7334 is also a valid IPv6 address(Omit leading zeros
and using upper cases).

However, we don't replace a consecutive group of zero value with a single empty group using two consecutive colons (::)
to pursue simplicity. For example, 2001:0db8:85a3::8A2E:0370:7334 is an invalid IPv6 address.

Besides, extra leading zeros in the IPv6 is also invalid. For example, the address
02001:0db8:85a3:0000:0000:8a2e:0370:7334 is invalid.

Note: You may assume there is no extra space or special characters in the input string.

Example 1:
Input: "172.16.254.1"

Output: "IPv4"

Explanation: This is a valid IPv4 address, return "IPv4".

Example 2:
Input: "2001:0db8:85a3:0:0:8A2E:0370:7334"

Output: "IPv6"

Explanation: This is a valid IPv6 address, return "IPv6".

Example 3:
Input: "256.256.256.256"

Output: "Neither"

Explanation: This is neither a IPv4 address nor a IPv6 address.
*/

func validIPAddress(IP string) string {
	for _, c := range IP {
		if c == '.' {
			return isIPv4(IP)
		} else if c == ':' {
			return isIPv6(IP)
		}
	}

	return neither
}

const neither = "Neither"

func isIPv4(ip string) string {
	partsCount := 0
	idx := 0
	for idx < len(ip) {
		r := 0
		i := 0
		for ; i < 3 && idx+i < len(ip); i++ {
			c := ip[idx+i]
			if c >= '0' && c <= '9' {
				if i > 0 && r == 0 {
					return neither
				}
				r = r*10 + int(c-'0')
			} else if c == '.' {
				break
			} else {
				return neither
			}
		}
		idx += i

		if i == 0 || r > 255 || (idx == len(ip)-1 && ip[idx] == '.') || (idx < len(ip) && ip[idx] != '.') {
			return neither
		}
		idx++

		partsCount++
		if partsCount > 4 {
			return neither
		}
	}
	if partsCount != 4 {
		return neither
	}

	return "IPv4"
}

func isIPv6(ip string) string {
	partsCount := 0
	idx := 0
	for idx < len(ip) {
		i := 0
		for ; i < 4 && idx+i < len(ip); i++ {
			if ip[idx+i] == ':' {
				break
			}
		}

		r := hexToInt(ip[idx : idx+i])
		idx += i

		if i == 0 || r < 0 || (idx == len(ip)-1 && ip[idx] == ':') || (idx < len(ip) && ip[idx] != ':') {
			return neither
		}
		idx++

		partsCount++
		if partsCount > 8 {
			return neither
		}
	}
	if partsCount != 8 {
		return neither
	}

	return "IPv6"
}

func hexToInt(hex string) int {
	result := 0

	for i, c := range hex {
		if c >= '0' && c <= '9' {
			if i > 1 && result == 0 {
				return -1
			}
			result <<= 4
			result += int(c - '0')
		} else if c >= 'A' && c <= 'F' {
			result <<= 4
			result += 10 + int(c-'A')
		} else if c >= 'a' && c <= 'f' {
			result <<= 4
			result += 10 + int(c-'a')
		} else {
			return -1
		}
	}

	return result
}
