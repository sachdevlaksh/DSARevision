/*
LeetCode Problem #1518: Water Bottles
Difficulty: Easy

There are numBottles water bottles that are initially full of water. You can exchange numExchange empty water bottles from the market with one full water bottle.
*/

package LC1518

func NumWaterBottles(numBottles int, numExchange int) int {
	out := numBottles
	for numBottles/numExchange > 0 {
		out += numBottles / numExchange
		numBottles = numBottles / numExchange + numBottles % numExchange
	}
	return out
}
