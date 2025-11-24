package LC3147

import "math"

func MaximumEnergy(energy []int, k int) int {
    n := len(energy)
    maxEnergy := math.MinInt

    // Iterate over all possible starting points in the last k positions
    for i := n - k; i < n; i++ {
        currentEnergy := 0
        for j := i; j >= 0; j -= k {
            currentEnergy += energy[j]
            if currentEnergy > maxEnergy {
                maxEnergy = currentEnergy
            }
        }
    }

    return maxEnergy
}