package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main2() {
	//subArraysWithSum([]int{0, 0, 0, 0, 0}, 0)
	//a1 := []int{1, 2}
	//a2 := []int{3, 2, 2, 1}
	//a3 := []int{3, 5, 3, 4}
	//a4 := []int{3, 3, 5}
	////a4 := []int{4, 5}
	////a5 := []int{12, 16}
	//a6 := [][]int{a1, a2, a3}
	//insert(a6, []int{4, 8})
	//findMinArrowShots(a6)
	//1-2, 3-5, 6-7, 8-10, 12-16
	//1-2,3-8,6-8,8-10,12-16
	//mergeInBetween(linkLi(), 3, 4, linkLi2())
	//timeRequiredToBuy([]int{2, 3, 2}, 2)
	//countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1})
	//longestIdealString("abcd", 4)
	//a := [][]int{
	//	{5, 2},
	//	{5, 4},
	//	{10, 3},
	//	{20, 1},
	//}
	//validPath(6, a, 0, 5)
	//minFallingPathSum(a6)
	//numIslands(a6)
	//reversePrefix("abcdefd", 'd')
	//fmt.Println(findMaxK([]int{14, 33, 40, -33, 8, -24, -42, 30, -18, -34}))
	//for i, in := range a6 {
	//	fmt.Println(numRescueBoats(in, a4[i])) // 1,2,4,5,
	//}
	//appendCharacters("coaching", "coding")
	//removeNodes(linkLi())

	//heightChecker([]int{1, 1, 4, 2, 1, 3})
	//minMovesToSeat([]int{2, 2, 6, 6}, []int{1, 3, 2, 6})
	//judgeSquareSum(4)
	//maxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2)
	//fmt.Println(maxHeightOfTriangle(10, 1))
	//fmt.Println(threeConsecutiveOdds([]int{1, 3}))
	//fmt.Println(intersect([]int{1, 2}, []int{1, 1}))
	//fmt.Println(minDifference([]int{1, 5, 0, 10, 24}))
	//fmt.Println(mergeNodes(linkLi()))
	//fmt.Println(nodesBetweenCriticalPoints(linkLi()))
	//fmt.Println(getEncryptedString("aaa", 1))
	//fmt.Println(numWaterBottles(9, 3))
	//fmt.Println(numberOfAlternatingGroups2([]int{0, 1, 0, 1, 0}, 3))
	//fmt.Println(findTheWinner(5, 2))
	//fmt.Println(averageWaitingTime(a))
	//fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	//fmt.Println(getSmallestString("20"))
	//fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	//	fmt.Println(nonSpecialCount(4, 16))
	//fmt.Println(canAliceWin([]int{5, 5, 5, 25}))
	//	n := 5
	//
	//	x := 10
	//	fmt.Println(minEnd(n, x)) // Output: 6
	//fmt.Println(makeFancyString("leeetcode"))
	//fmt.Print(isCircularSentence("MuFoevIXCZzrpXeRmTssj lYSW U jM"))
	//fmt.Println(rotateString("defdefdefabcabcdef", "defdefabcabcdefdef"))
	//	fmt.Println(compressedString("aaaaaaaaay"))
	//	fmt.Println(canSortArray([]int{3, 16, 8, 4, 2}))
	//	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 2))
	//	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	//	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	//	fmt.Println(isValidSudoku([][]byte{
	//		{'7','.','.','.','4','.','.','.','.'},
	//		{'.','.','.','8','6','5','.','.','.'},
	//		{'.','1','.','2','.','.','.','.','.'},
	//		{'.','.','.','.','.','9','.','.','.'},
	//		{'.','.','.','.','5','.','5','.','.'},
	//		{'.','.','.','.','.','.','.','.','.'},
	//		{'.','.','.','.','.','.','2','.','.'},
	//		{'.','.','.','.','.','.','.','.','.'},
	//		{'.','.','.','.','.','.','.','.','.'}}))
	//	fmt.Println(hIndex([]int{100}))
	//	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
	//	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(characterReplacement("AABABBA", 1))
}

func minEnd(n int, x int) int64 {
	n--
	ans := int64(x)
	for i := 0; i < 31; i++ {
		if (x>>i)&1 == 0 {
			ans |= int64((n & 1) << i)
			n >>= 1
		}
	}
	ans |= int64(n) << 31
	return ans
}
func subArraysWithSum(nums []int, goal int) int {
	var count int
	for length := 0; length < len(nums); length++ {
		count += findSubArray(nums, goal, 0, length)
	}
	return count

}

func findSubArray(nums []int, goal, length, count int) int {
	for i := 0; i < len(nums)-length; i++ {
		sum := findSum(nums, i, i+length)
		if sum == goal {
			count++
		}
	}
	return count

}

func findSum(nums []int, i, j int) int {
	sum := 0
	for i <= j {
		sum += nums[i]
		i++
	}
	return sum
}

func numSubarraysWithSum(nums []int, goal int) int {
	sumWithCount := make(map[int]int)
	sumWithCount[0] = 1
	count := 0
	sum := 0
	for _, num := range nums {
		sum += num
		val, _ := sumWithCount[sum-goal]
		count += val
		sumWithCount[sum] += 1
	}

	return count
}

//func productExceptSelf(nums []int) []int {
//
//	res := make([]int, len(nums))
//	res[0], res[len(nums)-1] = 1, 1
//	for i := 1; i < len(nums)-1; i++ {
//		res[i] = res[i-1] * nums[i-1]
//	}
//
//	rightProduct := 1
//	for i := len(nums) - 2; i >= 0; i-- {
//		rightProduct *= nums[i+1]
//		res[i] *= rightProduct
//	}
//
//	return res
//
//}

func insert(intervals [][]int, newInterval []int) [][]int {
	retArr := make([][]int, 0)
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			retArr = append(retArr, interval)
		} else if newInterval[1] < interval[0] {
			retArr = append(retArr, newInterval)
			newInterval = interval
		} else {
			newInterval[0] = int(math.Min(float64(newInterval[0]), float64(interval[0])))
			newInterval[1] = int(math.Max(float64(newInterval[1]), float64(interval[1])))
		}
	}
	retArr = append(retArr, newInterval)
	sort.SliceStable(retArr, func(i, j int) bool {
		return retArr[i][0] < retArr[j][0]
	})
	sort.SliceStable(retArr, func(i, j int) bool {
		return retArr[0][i] < retArr[0][j]
	})

	return retArr
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	count := 1
	lastMaxValue := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][1] <= lastMaxValue {
			continue
		}
		if lastMaxValue < points[i][0] {
			count++
			lastMaxValue = points[i][1]
		}
	}
	return count

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	count := 0
	res := new(ListNode)
	for list1.Next != nil {
		res.Val = list1.Val
		res.Next = list1.Next
		count++
		list1 = list1.Next
		if count == a {
			res.Next = list2
			break
		}
	}
	return res

}

func timeRequiredToBuy(tickets []int, k int) int {
	sum := 0
	num := tickets[k]
	sort.Sort(sort.Reverse(sort.IntSlice(tickets)))

	for i := len(tickets) - 1; i >= 0; i-- {
		if tickets[i] <= num {
			sum = sum + (len(tickets)-i)*tickets[i]
		}
	}
	return sum

}

func countStudents(students []int, sandwiches []int) int {

	count := make([]int, 2)
	for _, stud := range students {
		count[stud]++
	}
	for _, sand := range sandwiches {
		if count[sand] == 0 {
			return count[1-sand]
		} else {
			count[sand]--
		}
	}

	return 0
}

func validPath(n int, edges [][]int, start int, end int) bool {
	m := make(map[int][]int)

	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}

	alreadyDid := make(map[int]bool) // preventing infinite loops
	stack := []int{start}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop == end {
			return true
		}

		if !alreadyDid[pop] {
			alreadyDid[pop] = true
			stack = append(stack, m[pop]...)
		}
	}

	return false
}

func longestIdealString(s string, k int) int {
	var res int = 0

	var dp = [150]int{0}

	for _, c := range s {
		var pos = int(c)
		for i := pos - k; i <= pos+k; i++ {
			dp[pos] = int(math.Max(float64(dp[i]), float64(dp[pos])))
		}
		dp[pos]++
		res = int(math.Max(float64(dp[pos]), float64(res)))
	}

	return res
}

func minFallingPathSum(grid [][]int) int {
	mapS := make(map[int][]int)
	for i, v1 := range grid {
		li := []int{v1[i]}
		mapS[i] = li
		for j, v2 := range v1 {
			if i != j {
				mapS[i] = append(mapS[i], mapS[i][i]+v2)
			}
		}
	}
	return 13
}

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := make([][]int, n)
	indegree := make([]int, n)
	var ans []int

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
		indegree[e[0]]++
		indegree[e[1]]++
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 1 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		s := len(q)
		ans = []int{}
		for i := 0; i < s; i++ {
			pop := q[0]
			q = q[1:]
			ans = append(ans, pop)
			for _, child := range graph[pop] {
				indegree[child]--
				if indegree[child] == 1 {
					//push to Q
					q = append(q, child)
				}
			}
		}
	}

	if n == 1 {
		ans = append(ans, 0)
	}
	return ans
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	gr := make([][]int, n)
	q := make([][]int, 0)
	count := 0
	for i := 0; i < n; i++ {
		gr[i] = make([]int, m)
	}

	for ri, ro := range grid {
		for ci, co := range ro {
			if gr[ri][ci] != 1 && co == '1' {
				gr[ri][ci] = 1
				count++
				q = append(q, []int{ri, ci})
				for len(q) > 0 {
					pop := q[0]
					q = q[1:]
					for i := pop[0] - 1; i < n && i <= pop[0]+1; i++ {
						j := pop[1]
						if i < 0 {
							continue
						}

						if gr[i][j] != 1 && grid[i][j] == '1' {
							gr[i][j] = 1
							//fmt.Println("appending in I Loop : : ", i, j)
							q = append(q, []int{i, j})
						}
					}
					for j := pop[1] - 1; j < m && j <= pop[1]+1; j++ {
						i := pop[0]
						if j < 0 {
							continue
						}
						if gr[i][j] != 1 && grid[i][j] == '1' {
							gr[i][j] = 1
							//fmt.Println("appending in J Loop : : ", i, j)
							q = append(q, []int{i, j})
						}
					}
				}
			}
		}
	}
	//fmt.Println(count)
	return count
}

func reversePrefix(word string, ch byte) string {
	index := -1
	for i, c := range word {
		if string(ch) == string(c) {
			index = i
			break
		}
	}
	s := ""
	for i := index; i >= 0; i-- {
		s += string(word[i])
	}
	s = s + word[index+1:]
	return s
}

func findMaxK(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	if nums[len(nums)-1] < 0 {
		return -1
	}
	sort.Ints(nums)

	for i := 0; nums[i] < 0; i++ {
		find := nums[i]
		for j := len(nums) - 1; nums[j] > 0; j-- {
			if nums[j] == -1*find {
				return nums[j]
			}
		}
	}
	return -1
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	count := 0
	start := 0
	end := len(people) - 1

	for start <= end {
		if people[start]+people[end] <= limit {
			count++
			start++
			end--
		} else {
			count++
			end--
		}
	}

	return count
}

func deleteNode(node *ListNode) {

	for node.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
		if node.Next == nil {
			node = nil
		}
	}
}

func removeNodes(head *ListNode) *ListNode {
	stck := make([]int, 0)
	for head != nil {
		stck = append(stck, head.Val)

		head = head.Next
	}
	maxV := -1
	result := new(ListNode)
	result = nil
	for i := len(stck) - 1; i >= 0; i-- {
		if stck[i] >= maxV {
			temp := new(ListNode)
			temp.Val = stck[i]
			temp.Next = result
			result = temp
			maxV = stck[i]
		}

	}
	return result
}

func doubleIt(head *ListNode) *ListNode {
	stck := make([]int, 0)
	for head != nil {
		stck = append(stck, head.Val)

		head = head.Next
	}
	result := new(ListNode)
	car := 0
	result = nil
	for i := len(stck) - 1; i >= 0; i-- {
		val := 2 * stck[i]
		if val+car >= 10 {
			car = 1
		}
		temp := new(ListNode)
		temp.Val = val % 10
		temp.Next = result
		result = temp
	}
	stck2 := make([]int, 0)
	for result != nil {
		stck = append(stck2, result.Val)
		result = result.Next
	}
	return result
}

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l] = s[l] + s[r]
		s[r] = s[l] - s[r]
		s[l] = s[l] - s[r]
		l++
		r--
	}
}

func appendCharacters(s string, t string) int {
	qu := make([]byte, 0)
	for _, si := range t {
		qu = append(qu, byte(si))
	}
	for _, b := range s {
		if len(qu) > 0 && byte(b) == qu[0] {
			length := len(qu)
			qu = qu[1:length]
		}
	}
	return len(qu)

}

//func commonChars(words []string) []string {
//	av := make([]map[rune]int, 0)
//
//	for i, v := range words {
//		for _, r := range v {
//			av[i] = map[rune]int{r: av[i][r] + 1}
//		}
//	}
//
//}

func heightChecker(heights []int) int {
	cph := make([]int, len(heights))
	copy(cph, heights)

	sort.Ints(cph)
	count := 0
	for i, v := range heights {
		if cph[i] != v {
			count++
		}
		i++
	}
	return count
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	mape := make(map[int]int)
	for a, v := range arr2 {
		mape[v] = a
	}

	sort.Slice(arr1, func(i, j int) bool {
		rankI := rank(mape, arr1[i])
		rankJ := rank(mape, arr1[j])

		if rankI == rankJ {
			return arr1[i] < arr1[j]
		}
		return rankI < rankJ
	})

	return arr1
}

func rank(mape map[int]int, i int) int {
	num, ok := mape[i]
	if ok {
		return num
	}
	return len(mape)

}

func minMovesToSeat(seats []int, students []int) int {
	sliceSe := make([]int, slices.Max(seats)+1)
	sliceSt := make([]int, slices.Max(students)+1)

	for i, _ := range seats {
		sliceSe[seats[i]]++
		sliceSt[students[i]]++
	}

	p1 := 0
	p2 := 0
	count := 0
	remain := len(students)
	for remain > 0 {

		if sliceSe[p1] == 0 {
			p1 += 1
		}
		if sliceSt[p2] == 0 {
			p2 += 1
		}
		if sliceSe[p1] != 0 && sliceSt[p2] != 0 {
			count += int(math.Abs(float64(p1 - p2)))
			sliceSe[p1] -= 1
			sliceSt[p2] -= 1
			remain--
		}
	}
	return count
}

func canBePlaced(position []int, cows int, minMax int) bool {
	placed := 1
	start := position[0]
	for i := 0; i < len(position); i++ {
		if position[i]-start >= minMax {
			placed++
			start = position[i]
		}
		if placed >= cows {
			return true
		}
	}
	return false
}
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	length := len(position) - 1
	low, high := 1, position[length]-position[0]
	for low <= high {
		mid := low + (high-low)/2
		if canBePlaced(position, m, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return high

}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	iSum := 0
	for i, v := range grumpy {
		if v == 0 {
			iSum += customers[i]
		}
	}
	gSum := 0
	maxSum := 0
	for i := 0; i < len(grumpy); i++ {
		if i < minutes {
			if grumpy[i] == 1 {
				gSum += customers[i]
			}
		} else {
			if grumpy[i] == 1 {
				gSum += customers[i]
			}
			if grumpy[i-minutes] == 1 {
				gSum -= customers[i-minutes]
			}
		}
		maxSum = max(gSum, maxSum)
	}
	return maxSum + iSum
}

func findCenter(edges [][]int) int {
	mapA := make(map[int][]int)

	for _, vl := range edges {
		mapA[vl[0]] = append(mapA[vl[0]], vl[1])
		mapA[vl[1]] = append(mapA[vl[1]], vl[0])
	}
	for k, v := range mapA {
		if len(v) == len(edges) {
			return k
		}
	}
	return -1
}

func subarraySum(nums []int, k int) int {
	pam := make(map[int]int)
	pam[0] = 1
	preSum := 0
	count := 0
	for _, v := range nums {
		preSum += v
		if val, exist := pam[preSum-k]; exist {
			count = count + val
		}
		pam[preSum] = 1 + pam[preSum]
	}
	return count
}

func maximumImportance(n int, roads [][]int) int64 {
	list := make([]int64, n)
	for _, v := range roads {
		list[v[0]]++
		list[v[1]]++
	}
	sum := int64(0)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	for i := int64(1); i <= int64(n); i++ {
		sum += list[i] * i
	}
	return sum
}

func getAncestors(n int, edges [][]int) [][]int {
	pam := make([][]int, n)
	for _, v := range edges {
		pam[v[1]] = append(pam[v[1]], v[0])
	}
	ans := make([][]int, n)
	for k, mlv := range pam {
		q := mlv
		vis := make([]bool, n)
		for len(q) > 0 {
			i := q[0]
			q = q[1:]
			if !vis[i] {
				vis[i] = true
				ans[k] = append(ans[k], i)
				q = append(q, pam[i]...)
			}
		}
	}
	for _, v := range ans {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}
	return ans
}

//func getAncestors(n int, edges [][]int) [][]int {
//	g := make([][]int, n)
//	for _, e := range edges {
//		g[e[0]] = append(g[e[0]], e[1])
//	}
//	ans := make([][]int, n)
//	bfs := func(s int) {
//		q := []int{s}
//		vis := make([]bool, n)
//		vis[s] = true
//		for len(q) > 0 {
//			i := q[0]
//			q = q[1:]
//			for _, j := range g[i] {
//				if !vis[j] {
//					vis[j] = true
//					q = append(q, j)
//					ans[j] = append(ans[j], s)
//				}
//			}
//		}
//	}
//	for i := 0; i < n; i++ {
//		bfs(i)
//	}
//	return ans
//}

func maxHeightOfTriangle(red int, blue int) int {

	count1 := maxTriangle(red, blue)
	count := maxTriangle(blue, red)
	return max(count, count1)
}

func maxTriangle(red int, blue int) int {
	count := 0
	flip := true
	for red >= 0 && blue >= 0 {
		if flip {
			if red-(count+1) >= 0 {
				count++
			}
			red -= count
			flip = false

		} else {
			if blue-(count+1) >= 0 {
				count++
			}
			blue -= count
			flip = true
		}
	}
	return count
}

func threeConsecutiveOdds(arr []int) bool {
	for i, v := range arr {
		arr[i] = v % 2
	}
	low := 0
	right := 0
	sum := 0
	for right < len(arr) {
		if right <= 2 {
			sum += arr[right]
			right++
			if sum == 3 {
				return true
			}
		} else {
			sum += arr[right]
			right++
			sum = max(0, sum-arr[low])
			low++
			if sum == 3 {
				return true
			}

		}
	}
	return false
}

func intersect(nums1 []int, nums2 []int) []int {
	mx := 0
	mx1 := 0
	for _, v := range nums1 {
		if v > mx {
			mx = v
		}
	}
	for _, v := range nums2 {
		if v > mx1 {
			mx1 = v
		}
	}
	dp1 := make([]int, mx+1)
	dp2 := make([]int, mx1+1)

	for _, v := range nums1 {
		dp1[v]++
	}
	for _, v := range nums2 {
		dp2[v]++
	}

	out := make([]int, 0)
	if mx < mx1 {
		for i, v := range dp1 {
			if v > 0 && dp2[i] > 0 {
				index := 0
				for index < min(v, dp2[i]) {
					out = append(out, i)
					index++
				}
			}
		}
	} else {
		for i, v := range dp2 {
			if v > 0 && dp1[i] > 0 {
				index := 0
				for index < min(v, dp1[i]) {
					out = append(out, i)
					index++
				}
			}
		}
	}
	return out

}

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	l := len(nums)
	minDiff := int(^uint(0) >> 1)
	for i := 0; i < 4; i++ {
		diff := nums[l-4+i] - nums[i]
		if minDiff > diff {
			minDiff = diff
		}
	}

	return minDiff
}

//func minDifference(nums []int) int {
//	if len(nums) <= 4 {
//		return 0
//	}
//	sort.Ints(nums)
//	minDiff := int(^uint(0) >> 1)
//	for left := 0; left <= 3; left++ {
//		diff := nums[len(nums)-4+left] - nums[left]
//		if diff < minDiff {
//			minDiff = diff
//		}
//	}
//	return minDiff
//}

func mergeNodes(head *ListNode) *ListNode {

	list := &ListNode{}
	curr := list
	for head != nil {
		if head.Val == 0 && head.Next != nil {
			curr.Next = &ListNode{Val: 0}
			curr = curr.Next
		} else {
			curr.Val += head.Val
		}
		head = head.Next
	}
	return list.Next
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	li := make([]int, 0)
	prev := head.Val
	head = head.Next
	count := 1
	mini := int(^uint(0) >> 1)
	for head.Next != nil {
		if head.Val < prev && head.Val < head.Next.Val || head.Val > prev && head.Val > head.Next.Val {
			li = append(li, count)
			if len(li) >= 2 {
				if mini > li[len(li)-1]-li[len(li)-2] {
					mini = li[len(li)-1] - li[len(li)-2]
				}
			}
		}
		prev = head.Val
		head = head.Next
		count++
	}
	if len(li) < 2 {
		return []int{-1, -1}
	} else if len(li) == 2 {
		return []int{li[1] - li[0], li[1] - li[0]}
	} else {
		return []int{mini, li[len(li)-1] - li[0]}
	}

}

func numberOfAlternatingGroups(colors []int) int {
	count := 0
	if len(colors) >= 3 {
		colors = append(colors, colors[0])
		colors = append(colors, colors[1])
		r := 1
		for r < len(colors)-1 {
			if (colors[r] != colors[r+1]) && (colors[r] != colors[r-1]) {
				count++
			}
			r++
		}
	}
	return count
}
func numberOfAlternatingGroups2(colors []int, k int) int {
	count := 1
	colors = append(colors, colors[:k-1]...)
	l := 1
	res := 0
	for l < len(colors) {
		if colors[l] != colors[l-1] {
			count++
		} else {
			count = 1
		}
		if count >= k {
			res++
		}
		l++
	}

	return res
}

func getEncryptedString(s string, k int) string {
	if len(s) <= 1 {
		return s
	}
	newS := []rune(s + s)
	runeS := []rune(s)
	for i, _ := range s {
		runeS[i] = newS[i+k%len(s)]
	}
	var sout string

	for _, v := range runeS {
		sout += string(v)
	}
	return sout
}

func numWaterBottles(numBottles int, numExchange int) int {
	sum := numBottles
	for numBottles >= numExchange {
		sum += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return sum
}

func findTheWinner(n int, k int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i + 1
	}
	i := 0
	for len(dp) > 1 {
		i = (i + k - 1) % len(dp)
		dp = append(dp[:i], dp[i+1:]...)
	}
	return dp[0]
}

func averageWaitingTime(cus [][]int) float64 {
	prevSum := float64(cus[0][0] + cus[0][1])
	waitTime := 0.0
	waitTime += float64(cus[0][1])
	i := 1

	for i < len(cus) {
		if prevSum >= float64(cus[i][0]) {
			waitTime += float64(prevSum + float64(cus[i][1]-cus[i][0]))
			prevSum += float64(cus[i][1])
		} else {
			waitTime += float64(cus[i][1])
			prevSum = float64(cus[i][0] + cus[i][1])
		}
		i++

	}

	return waitTime / float64(len(cus))

}

func minOperations(logs []string) int {
	count := 0
	for _, log := range logs {
		switch log {
		case "../":
			count--
		case "./":
		default:
			count++

		}
	}
	return count
}

func reverseParentheses(s string) string {

	startParenthIndexes := make([]int, 0)
	count := 0
	for parenthIndex, si := range s {
		ss := string(si)
		switch ss {
		case "(":
			startParenthIndexes = append(startParenthIndexes, parenthIndex)
		case ")":

			preIndex := startParenthIndexes[len(startParenthIndexes)-1] - 2*count
			sl := len(s)
			so := ""
			so += s[0:preIndex]
			so += reverseStr(s[preIndex+1 : parenthIndex])
			if parenthIndex < sl-1 {
				so += s[parenthIndex+1:]
			}
			s = so
			count++
			startParenthIndexes = startParenthIndexes[:len(startParenthIndexes)-1]
		default:
		}
	}
	return s
}

func reverseStr(str string) string {
	strRune := []rune(str)
	i := 0
	j := len(strRune) - 1
	outString := ""
	for i <= j {
		strRune[i], strRune[j] = strRune[j], strRune[i]
		i++
		j--
	}
	for _, v := range strRune {
		outString += string(v)
	}
	return outString
}

func getSmallestString(s string) string {
	prev, _ := strconv.Atoi(string(s[0]))
	i := 1
	for i < len(s) {
		curr, _ := strconv.Atoi(string(s[i]))
		if curr < prev && (int(s[i])%2 == prev%2) {
			so := s[0:i-1] + string(s[i]) + strconv.Itoa(prev)
			if i+1 < len(s) {
				so += s[i+1:]
			}
			s = so
			break
		}
		prev = curr
		i++
	}
	return s
}

func sortPeople(names []string, heights []int) []string {
	i, j := 0, len(heights)-1

	for j >= 0 {
		for i <= j {
			if heights[i] < heights[j] {
				heights[j], heights[i] = heights[i], heights[j]
				names[j], names[i] = names[i], names[j]
				i++
			}
		}
		i = 0
		j--
	}
	return names

}

func sortArray(nums []int) []int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func nonSpecialCount(l int, r int) int {
	mapDiv := make(map[int]interface{})
	rc := r - l + 1
	for l <= r {
		count := 1
		sqrtL := int(math.Sqrt(float64(l)))
		if _, ok := mapDiv[sqrtL]; ok {
			count = 2
			break
		} else {
			sqrtLL := int(math.Sqrt(float64(sqrtL)))
			if _, ok := mapDiv[sqrtLL]; ok {
				count = 2
				break
			} else {
				for i := 2; i <= sqrtL; i++ {
					if l%i == 0 {
						count++
						if l > i*i {
							count = 1
							break
						}
					}

					if count > 2 {
						break
					}
				}
			}
		}
		l++
		if count == 2 {
			mapDiv[l-1] = true
			rc--
		}
	}
	return rc
}

func canAliceWin(nums []int) bool {
	count1 := 0
	count := 0
	for _, num := range nums {
		if num/10 < 1 {
			count1 += num
		} else {
			count += num
		}
	}
	if count != count1 {
		return true
	}
	return false
}

func makeFancyString(s string) string {

	var sb strings.Builder
	for len(s) > 2 {
		if s[2] != s[1] || s[2] != s[0] {
			sb.WriteString(string(s[0]))

		}
		s = s[1:]
	}
	sb.WriteString(s)
	return sb.String()
}

func isCircularSentence(sentence string) bool {

	if sentence[0] != sentence[len(sentence)-1] || string(sentence[0]) == "" || string(sentence[len(sentence)-1]) == "" {
		return false
	}

	for i, s := range sentence {
		if string(s) == " " && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}

func rotateString(s string, goal string) bool {
	length := len(s)
	if len(goal) != len(s) {
		return false
	}
	for i := 0; i < length; i++ {
		if goal[i:] == s[:length-i] {
			if goal[:i] == s[length-i:] {
				return true
			}
		}
	}
	return false
}

func compressedString(word string) string {
	if len(word) == 1 {
		return "1" + word
	}
	var sb strings.Builder
	count := 1
	prev := word[0]

	for i := 1; i < len(word); i++ {
		if word[i] != prev {
			if count > 0 {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteString(string(word[i-1]))
			}
			count = 1
			prev = word[i]
		} else {
			count++
			if count == 9 {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteString(string(word[i]))
				count = 0
			}
		}
		if i == len(word)-1 && count > 0 {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteString(string(word[i]))
		}
	}

	return sb.String()
}

func canSortArray(nums []int) bool {
	maxAtt := len(nums) - 1
	att := 0
	for att < maxAtt {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				if getSetBit(nums[i+1]) == getSetBit(nums[i]) {
					nums[i], nums[i+1] = nums[i+1], nums[i]
					att = 0
				} else {
					return false
				}
			} else {
				att++
			}
		}
	}
	return true
}

func getSetBit(num int) int {
	count := 0
	for num > 0 {
		if num&1 != 0 {
			count++
		}
		num >>= 1
	}
	return count

}

func getMaximumXor(nums []int, maximumBit int) []int {
	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			out[i] = ^nums[j]
			out[i] = out[i]&(1<<maximumBit) - 1
		}
	}
	return out
}

func minimumSubarrayLength(nums []int, k int) int {
	sum := 0
	l, r := 0, 1
	sum += nums[l]
	for r < len(nums) {
		if sum-nums[l] > k {
			sum -= nums[l]
			l++
		} else if sum >= k && sum-nums[l] < k {
			return r - l
		} else if sum < k {
			sum += nums[r]
			r++
		}
	}
	return r - l + 1
}

type key [26]byte

func groupAnagrams(strs []string) (out [][]string) {
	outRes := make(map[key][]string)
	for _, stri := range strs {
		uniqueKey := getKeys(stri)
		outRes[uniqueKey] = append(outRes[uniqueKey], stri)
	}

	for _, res := range outRes {
		out = append(out, res)
	}

	return out
}

func getKeys(ipstr string) (outkey key) {
	for _, ips := range ipstr {
		outkey[ips-'a']++
	}

	return outkey
}

func topKFrequent(arr []int, k int) []int {
	out := make(map[int]int)
	for _, ipArrV := range arr {
		out[ipArrV]++

	}
	slic := make([][2]int, 0)
	for ke, v := range out {
		slic = append(slic, [2]int{v, ke})
	}

	sort.Slice(slic, func(i, j int) bool {
		return slic[i][0] > slic[j][0]
	})

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, slic[i][1])
	}

	return result
}

func productExceptSelf(nums []int) []int {
	var singleZero bool
	prod := 1
	for _, num := range nums {
		if num == 0 && !singleZero {
			singleZero = true
			continue
		}
		prod = prod * num
	}

	for numI, num := range nums {
		if singleZero && num != 0 {
			nums[numI] = 0
		} else if singleZero && num == 0 {
			nums[numI] = prod
		} else if !singleZero {
			nums[numI] = prod / nums[numI]
		}
	}
	return nums

}

func isValidSudoku(board [][]byte) bool {
	var li2 []byte
	var li3 []byte
	var li1 []byte
	for i := 0; i < len(board); i++ {
		li1 = []byte{}
		for p := 0; p < len(board); p++ {
			if board[i][p] != '.' {
				li1 = append(li1, board[i][p])
			}
		}

		for j := 0; j < len(board); j++ {
			li2 = []byte{}
			for p := 0; p < len(board); p++ {
				if board[p][j] != '.' {
					li2 = append(li2, board[p][j])
				}
			}

			if i%3 == 0 && j%3 == 0 {
				li3 = []byte{}
				for k := i; k <= i+2; k++ {
					for l := j; l <= j+2; l++ {
						if board[k][l] != '.' {
							li3 = append(li3, board[k][l])
						}
					}
				}
			}
			if isDuplicated(li1) || isDuplicated(li2) || isDuplicated(li3) {
				return false

			}

		}
	}
	return true
}

func isDuplicated(firstList []byte) bool {
	mapp := make(map[byte]interface{})
	for _, val := range firstList {
		if _, ok := mapp[val]; ok {
			return true
		} else {
			mapp[val] = true
		}
	}
	return false
}

func hIndex(cits []int) int {
	if len(cits) == 1 {
		if cits[0] == 0 {
			return 0
		} else {
			return 1
		}
	}
	citations := make([]int, 0)
	for _, cite := range cits {
		if cite != 0 {
			citations = append(citations, cite)
		}
	}
	size := len(citations)
	sort.Ints(citations)

	lastOcc := 0
	hIndex := make([]int, size)
	hIndex[0] = size
	lastOcc = size
	for i := 1; i < len(citations); i++ {
		if citations[i] == citations[i-1] {
			hIndex[i] = lastOcc
		} else {
			hIndex[i] = size - i
		}
		lastOcc = hIndex[i]
	}
	for i := size - 1; i >= 0; i-- {
		if citations[i] <= hIndex[i] {
			return citations[i]
		}
	}

	return int(math.Max(float64(size), 0))
}

func longestConsecutive(nums []int) int {
	maxLength := 1
	hm := make(map[int]interface{})
	for _, num := range nums {
		hm[num] = true
	}
	for _, num := range nums {
		if _, ok := hm[num]; ok {
			length := 1
			for {
				if _, ok2 := hm[num+length]; ok2 {
					length++
				} else {
					break
				}
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

func isPalindrome(s string) bool {

	l, r := 0, len(s)-1

	for l < r {

		for l < r && !isAlphaNumeric(s[l]) {
			l++
		}
		for l < r && !isAlphaNumeric(s[r]) {
			r--
		}

		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}

	return true

}

func isAlphaNumeric(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '1' && c <= '9') {
		return true
	}
	return false
}

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] > target {
			r--
		}
	}
	return []int{0, 0}
}

func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	sort.Ints(nums)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := 0, len(nums)-1
		for l < r && i != l && i != r {
			if nums[l]+nums[r]+num == 0 {
				out = append(out, []int{nums[l], num, nums[r]})
			} else if nums[l]+nums[r]+num > 0 {
				r--
			} else if nums[l]+nums[r]+num < 0 {
				l++
			}
			l++
			r--
			for l < r && nums[l] == nums[l-1] {
				l++
			}
		}

	}

	return out

}

func maxArea(height []int) int {

	l, r := 0, len(height)
	maxAr := 1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		if area > maxAr {
			maxAr = area

		}
	}
	return maxAr
}

func maxProfit(prices []int) int {
	low := 0
	high := low + 1

	diff := 0
	for low < high && high < len(prices) {
		maxDiff := prices[high] - prices[low]
		if diff < maxDiff {
			diff = maxDiff
		}
		high++
		if high == len(prices) {
			low++
			high = low + 1
		}

	}
	return diff
}

func characterReplacement(s string, k int) int {
	hmap := make(map[byte]int)
	l, maxF, res := 0, 0, 0

	for r := 0; r < len(s); r++ {
		hmap[s[r]]++
		if hmap[s[r]] > maxF {
			maxF = hmap[s[r]]
		}
		for r-l+1-maxF > k {
			hmap[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
func AddGigasecond(t time.Time) time.Time {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	t.Add(time.Second * 1000)

	return t
}

func CollatzConjecture(n int) (int, error) {
	count := 0
	if n == 0 || n < 0 {
		return 0, errors.New("error")
	}
	for n > 1 {
		if n%2 == 1 {
			n = (3 * n) + 1
		} else {
			n = n / 2
		}
		count++
	}

	return count, nil
}

func Score(word string) int {
	mapp := make(map[rune]int)
	mapp['A'] = 1
	mapp['E'] = 1
	mapp['I'] = 1
	mapp['O'] = 1
	mapp['U'] = 1
	mapp['L'] = 1
	mapp['N'] = 1
	mapp['R'] = 1
	mapp['S'] = 1
	mapp['T'] = 1
	mapp['D'] = 2
	mapp['G'] = 2
	mapp['B'] = 3
	mapp['C'] = 3
	mapp['M'] = 3
	mapp['P'] = 3
	mapp['F'] = 4
	mapp['H'] = 4
	mapp['V'] = 4
	mapp['W'] = 4
	mapp['Y'] = 4
	mapp['K'] = 5
	mapp['J'] = 8
	mapp['X'] = 8
	mapp['Q'] = 10
	mapp['Z'] = 10
	code := 0
	word = strings.ToUpper(word)
	for _, ai := range word {
		code += mapp[rune(ai)]
	}

	return code
}

func fun(m map[int][]int , cost []int , curNode int , ans *int) int64 {
    pathSum := int64(cost[curNode])
    maxPathSum := int64(0)
    pathSumVals := []int64{}

    for _ , child := range m[curNode] {
        childPathSum := fun(m , cost , child , ans)
        pathSumVals = append(pathSumVals , childPathSum)
        if childPathSum > maxPathSum {
            maxPathSum = childPathSum
        }
    }

    for _ , val := range pathSumVals {
        if val != maxPathSum {
            *ans += 1
        }
    }

    return pathSum + maxPathSum
}
func minIncrease(n int, edges [][]int, cost []int) int {
    m := make(map[int][]int)

    for _ , edge := range edges {
        m[edge[0]] = append(m[edge[0]] , edge[1])
    }

    ans := 0

    fun(m , cost , 0 , &ans)

    return ans
}

func findCoins(numWays []int) []int {
	n := len(numWays)
	numWays = append([]int{1}, numWays...)
	myWays := make([]int, n+1)
	myWays[0] = 1
	res := []int{}

	for i := 1; i <= n; i++ {
		// If `myWays[x] == numWays[x]`, move on.
		if myWays[i] == numWays[i] {
			continue
		}

		// If `myWays[x] + 1 == numWays[x]` → add that value as a coin in our basket and update `myWays`, so `myWays[x...n]` accounts for ways with the new coin.
		if numWays[i]-myWays[i] == 1 {
			res = append(res, i)
			for j := i; j <= n; j++ {
				myWays[j] += myWays[j-i]
			}
		// If `myWays[x] + 1 < numWays[x]` → no solution. (*see below for why)
		} else {
			return []int{}
		}
	}
	return res
}