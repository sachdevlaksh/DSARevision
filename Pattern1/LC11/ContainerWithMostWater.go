package LC11

func maxArea(height []int) int {
    i := 0
    j:= len(height) -1
    maxArea := 0
    for i < j {
        maxArea = findMax(maxArea,findMin(height[i],height[j])*(j-i))
        if height[i]< height[j]{
            i++
        }else{
            j--
        }
    }
    return maxArea
}

func findMax( i, j  int) int{
    if i < j {
        return j
    }else{
        return i
    }
}

func findMin( i, j  int) int{
    if i < j {
        return i
    }else{
        return j
    }
}