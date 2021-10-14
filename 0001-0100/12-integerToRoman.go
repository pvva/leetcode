package main

func main() {
	for _, n := range []int{3, 4, 8, 9, 12, 18, 27, 49, 449, 949, 999, 1998, 2021} {
		println(n, intToRoman(n), intToRoman2(n))
	}
}

var intToRomanValues = []struct {
	v int
	r string
}{
	{v: 1000, r: "M"},
	{v: 500, r: "D"},
	{v: 100, r: "C"},
	{v: 50, r: "L"},
	{v: 10, r: "X"},
	{v: 5, r: "V"},
	{v: 1, r: "I"},
}

func getSpecialRomanNumber(v int) (string, int) {
	r := ""
	rv := v
	if v == 4 {
		r += "IV"
		rv = 0
	} else if v == 9 {
		r += "IX"
		rv = 0
	} else if v >= 40 && v <= 49 {
		r += "XL"
		rv = v - 40
	} else if v >= 90 && v <= 99 {
		r += "XC"
		rv = v - 90
	} else if v >= 400 && v <= 499 {
		r += "CD"
		rv = v - 400
	} else if v >= 900 && v <= 999 {
		r += "CM"
		rv = v - 900
	}
	return r, rv
}

func intToRoman(num int) string {
	r := ""

	for _, s := range intToRomanValues {
		if num == 0 {
			break
		}
		for {
			sn, n := getSpecialRomanNumber(num)
			if sn == "" {
				break
			}
			r += sn
			num = n
		}

		i := num / s.v
		if i == 0 {
			continue
		}
		d := i * s.v
		num -= d
		for ; i > 0; i-- {
			r += s.r
		}
	}

	return r
}

var intToRomanValues2 = []struct {
	v int
	r string
}{
	{v: 1000, r: "M"},
	{v: 900, r: "CM"},
	{v: 500, r: "D"},
	{v: 400, r: "CD"},
	{v: 100, r: "C"},
	{v: 90, r: "XC"},
	{v: 50, r: "L"},
	{v: 40, r: "XL"},
	{v: 10, r: "X"},
	{v: 9, r: "IX"},
	{v: 5, r: "V"},
	{v: 4, r: "IV"},
	{v: 1, r: "I"},
}

func intToRoman2(num int) string {
	r := ""

	for _, s := range intToRomanValues2 {
		for num >= s.v {
			num -= s.v
			r += s.r
		}
	}

	return r
}
