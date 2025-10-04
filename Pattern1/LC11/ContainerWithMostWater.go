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
