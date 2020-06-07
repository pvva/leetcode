package main

func main() {
	println(rotationalCipher("Zebra-493?", 3), "Cheud-726?")
	println(rotationalCipher("abcdefghijklmNOPQRSTUVWXYZ0123456789", 39), "nopqrstuvwxyzABCDEFGHIJKLM9012345678")
}

func rotationalCipher(input string, rotationFactor int) string {
	result := make([]byte, len(input))

	for i, c := range input {
		result[i] = getCipheredSymbol(byte(c), rotationFactor)
	}

	return string(result)
}

func getCipheredSymbol(c byte, rotation int) byte {
	base := byte(0)
	rng := 0
	if c >= '0' && c <= '9' {
		base = '0'
		rng = 10
	} else if c >= 'a' && c <= 'z' {
		base = 'a'
		rng = 26
	} else if c >= 'A' && c <= 'Z' {
		base = 'A'
		rng = 26
	}

	if base == 0 {
		return c
	}

	return byte((int(c-base)+rotation)%rng) + base
}
