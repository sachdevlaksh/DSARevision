/*
LeetCode Problem #11: Container With Most Water
Difficulty: Medium

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form a container, such that the container contains the most water.
*/

package LC11

func maxArea(height []int) int {
    i := 0
    j := len(height) - 1
    maxArea := 0
    for i < j {
        maxArea = max(maxArea, min(height[i], height[j])*(j-i))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return maxArea
}
