/*
LeetCode Problem #134: Gas Station
Difficulty: Medium

There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i]. You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to the (i + 1)th station.
*/

package LC134

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	if n == 0 {
		return -1
	}

	gasT := 0
	costT := 0
	for i := 0; i < n; i++ {
		gasT += gas[i]
		costT += cost[i]
	}
	if costT > gasT {
		return -1
	}

	tank := 0
	start := 0

	for i := 0; i < n; i++ {
		tank += gas[i] - cost[i]

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	return start
}
