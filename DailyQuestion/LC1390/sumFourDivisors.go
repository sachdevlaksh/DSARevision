package main

func sumFourDivisors(nums []int) int {

	sum := 0
	for _, n := range nums {
		totalSum := 1 + n
		totalDivisor := 2

		for i := 2; i*i <= n && totalDivisor <= 4; i++ {
			if n%i == 0 {
				j := n / i

				if i == j {
					totalDivisor += 1
					totalSum += i
				} else {
					totalDivisor += 2
					totalSum += i + j
				}
			}
		}

		if totalDivisor == 4 {
			sum += totalSum
		}

	}

	return sum
}
