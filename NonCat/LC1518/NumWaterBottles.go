package LC1518

func NumWaterBottles(numBottles int, numExchange int) int {
	out := numBottles
	for numBottles/numExchange > 0 {
		out += numBottles / numExchange
		numBottles = numBottles / numExchange + numBottles % numExchange
	}
	return out
}
