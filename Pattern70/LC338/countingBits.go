package LC338

func countBits(n int) []int {
    bits := make([]int, n+1)

    for i:= 0; i < n+1; i++{
        bits[i] = bits[i>>1] + i & 1
    }

    return bits[:n+1]
}