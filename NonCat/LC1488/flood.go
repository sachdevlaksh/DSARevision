package LC1488

import (
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	mp := make(map[int]int)
	zeroIndexes := []int{}

	for i := 0; i < n; i++ {
		if rains[i] == 0 {
			zeroIndexes = append(zeroIndexes, i)
			continue
		}

		if last, found := mp[rains[i]]; found {
			if len(zeroIndexes) == 0 {
				return []int{}
			}

			// Find the smallest element in zeroIndexes >= last
			pos := sort.Search(len(zeroIndexes), func(j int) bool { return zeroIndexes[j] > last })
			if pos == len(zeroIndexes) {
				return []int{}
			}

			ans[zeroIndexes[pos]] = rains[i]
			zeroIndexes = append(zeroIndexes[:pos], zeroIndexes[pos+1:]...)
		}

		ans[i] = -1
		mp[rains[i]] = i
	}

	for i := 0; i < n; i++ {
		if ans[i] == 0 {
			ans[i] = 1
		}
	}

	return ans
}

func AvoidFlood(rains []int) []int {
	zeroRainsIndex := make([]int, 0)
	pastRain := make(map[int]int)
	for i, rain := range rains {

		if rain == 0 {
			zeroRainsIndex = append(zeroRainsIndex, i)
			continue
		}
		rains[i] = -1
		if _, found := pastRain[rain]; found {
			rainNextZeroIndex := searchRain(pastRain[rain], &zeroRainsIndex)
			if rainNextZeroIndex == -1 {
				return []int{}
			}
			rains[rainNextZeroIndex] = rain
		}
		pastRain[rain] = i

	}

	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			rains[i] = 1
		}
	}
	return rains
}

func searchRain(rain int, allZeros *[]int) int {
	for i, v := range *allZeros {
		if v > rain {
			*allZeros = append((*allZeros)[:i], (*allZeros)[i+1:]...)
			return v

		}
	}

	return -1
}
