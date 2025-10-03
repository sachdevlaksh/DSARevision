package main

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {

	res := ""
	if numerator == 0 {
		return "0"
	}

	n := int64(numerator)
	d := int64(denominator)

	if (n < 0) != (d < 0) {
		res += "-"
	}

	if n < 0 {
		n = -n
	}
	if d < 0 {
		d = -d
	}
	res += strconv.FormatInt(n/d, 10)
	rem := n % d

	if rem == 0 {
		return res
	}

	res += "."
	remMap := make(map[int64]int)
	for rem != 0 {
		if pos, ok := remMap[rem]; ok {
			res = res[:pos] + "(" + res[pos:] + ")"
			break
		}
		remMap[rem] = len(res)
		rem *= 10
		res += strconv.FormatInt(rem/d, 10)
		rem %= d
	}
	return res
}
