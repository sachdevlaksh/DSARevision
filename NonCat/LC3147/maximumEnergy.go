/*
LeetCode Problem #3147: Taking Maximum Energy From The Mystic Dungeon
Difficulty: Medium

In a mystic dungeon, n magicians are standing in a line. Each magician has an attack power denoted by a[i]. Your task is to collectthe maximum energy possible.
*/

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