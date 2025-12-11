package LC45

func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    end := 0
    farthest := 0
    jumps := 0

    for i := 0; i < len(nums); i++ {

        if i+nums[i] > farthest {
            farthest = i + nums[i]
        }

        if i == end {
            jumps++
            end = farthest
        }

        if end >= len(nums)-1 {
            break
        }
    }

    return jumps
}
