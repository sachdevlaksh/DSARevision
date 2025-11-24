package LC56

import "sort"

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
        return nil
    }
	out := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0]{
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	i := 0;
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		for i+1 < len(intervals) && intervals[i+1][0] <= end{
			i++
			if intervals[i][1]>end{
				end = intervals[i][1]
			}
		}
		out = append(out, []int{start,end})
		i++
	}
	return out
}