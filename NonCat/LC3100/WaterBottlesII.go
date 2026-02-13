/*
LeetCode Problem #3100: Water Bottles II
Difficulty: Medium

Alice has n full water bottles. In one operation, she can empty x bottles and fill one bottle using the water from those x bottles.
*/

package LC3100

func MaxBottlesDrunk(numBottles int, numExchange int) int {

	out := numBottles
	emptyBottles :=  numBottles
	extraBottle := 0
	for emptyBottles>= numExchange {
		extraBottle++
		emptyBottles -= numExchange
		emptyBottles++
		numExchange++
	}

	return out + extraBottle
}