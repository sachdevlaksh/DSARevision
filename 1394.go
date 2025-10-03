package main

import (
	"fmt"
	"strconv"
)

type vari struct {
	Variant            string `json:"variant,omitempty"`
	Size1              string `json:"size1,omitempty"`
	Size2              string `json:"size2,omitempty"`
	SizeSelectorLabel1 string `json:"sizeSelectorLabel1,omitempty"`
	SizeSelectorLabel2 string `json:"sizeSelectorLabel2,omitempty"`
	SizeCode           int    `json:"sizeCode,omitempty"`
	SizeSortSequence   int    `json:"sizeSortSequence,omitempty"`
	SizeSortSequence2  int    `json:"sizeSortSequence2,omitempty"`
}

//func main() {
//	//findLucky([]int{2, 2, 3, 4})
//	//linkLi()
//	//addTwoNumbers(l1, l2)
//	//fmt.Print(lengthOfLongestSubstring("abcabcbb"))
//	//fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
//	//isValidStack("([{}])")
//	//generateParenthesis(3)
//	//nextPermutation([]int{1, 3, 2})
//	//fmt.Print(combinationSum([]int{2, 3, 5}, 8))
//	//fmt.Print(Permutation([]int{1, 2, 3}))
//	//fmt.Print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
//	//fmt.Print(twoSum([]int{0, 4, 3, 0}, 0))
//	//fmt.Println(myPow(2, 10))
//	//fmt.Println(firstUniqChar("av"))
//	//li := linkLi()
//	//fmt.Println(deleteDuplicates(li))
//	//fmt.Println(minAddToMakeValid("()))"))
//
//	filePath := "./2.json"
//	fmt.Printf("// reading file %s\n", filePath)
//	file, err1 := ioutil.ReadFile(filePath)
//	if err1 != nil {
//		fmt.Printf("// error while reading file %s\n", filePath)
//		fmt.Printf("File error: %v\n", err1)
//		os.Exit(1)
//	}
//	fmt.Println("// defining array of struct shipObject")
//	var variant vari
//
//	err2 := json.Unmarshal(file, &variant)
//	if err2 != nil {
//		fmt.Println("error:", err2)
//		os.Exit(1)
//	}
//	//for _, ship := range variants {
//	fmt.Println(variant.variant)
//	//}
//
//}

func generateParenthesis(n int) []string {
	var res []string

	backtrack(&res, []byte{}, n, 0, 0)

	return res
}

func backtrack(res *[]string, s []byte, n int, left int, right int) {
	if len(s) == n*2 {
		*res = append(*res, string(s))
	}

	if left < n {
		backtrack(res, append(s, '('), n, left+1, right)
	}

	if right < left {
		backtrack(res, append(s, ')'), n, left, right+1)
	}
}
func isValidStack(s string) bool {
	if len(s) <= 1 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			lastIndex := len(stack) - 1
			if lastIndex == -1 {
				return false
			}
			lastOpenParenthese := stack[lastIndex]
			if v == m[lastOpenParenthese] {
				stack = stack[:lastIndex]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func linkLi() *ListNode {
	l1 := new(ListNode)
	l1.Val = 5
	l1.Next = &ListNode{
		3,
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					5,
					&ListNode{
						1,
						&ListNode{
							2,
							nil},
					},
				},
			},
		},
	}
	l2 := new(ListNode)
	l2.Val = 10000
	l2.Next = &ListNode{
		10001,
		&ListNode{
			100002,
			nil,
			//&ListNode{
			//	6,
			//	&ListNode{
			//		2,
			//		&ListNode{
			//			9,
			//			nil,
			//		},
			//	},
			//},
		},
	}
	return l1
}
func linkLi2() *ListNode {
	l2 := new(ListNode)
	l2.Val = 10000
	l2.Next = &ListNode{
		10001,
		&ListNode{
			100002,
			nil,
			//&ListNode{
			//	6,
			//	&ListNode{
			//		2,
			//		&ListNode{
			//			9,
			//			nil,
			//		},
			//	},
			//},
		},
	}
	return l2
}
func findLucky(arr []int) int {
	occurenceCount := make(map[string]int)
	for _, arrV := range arr {
		if _, exist := occurenceCount[string(arrV)]; !exist {
			occurenceCount[string(rune(arrV))] = 1
		} else {
			occurenceCount[string(arrV)] = occurenceCount[string(arrV)] + 1
		}
	}
	minValue := 99999
	for k, v := range occurenceCount {
		fmt.Print(v)
		key, _ := strconv.Atoi(k)
		if v == key {
			if key < minValue {
				minValue = key
			}
		}
	}
	return minValue
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	cur := new(ListNode)
	ret := cur
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			sum % 10,
			nil,
		}
		cur = cur.Next
		sum = sum / 10
	}

	if sum > 0 {
		cur.Next = &ListNode{sum, nil}
	}
	return ret.Next
}

//func lengthOfLongestSubstring(s string) int {
//	mapHelper := make(map[string]int)
//	start := 0
//	end := 0
//	maxLength := 0
//	for end = 0; end < len(s); end++ {
//		if prevEndIndex, ok := mapHelper[string(s[end])]; ok {
//			start = max(start, prevEndIndex+1)
//		}
//		mapHelper[string(s[end])] = end
//		maxLength = max(maxLength, end-start+1)
//	}
//	return maxLength
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if nums[left] > nums[mid] && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] < nums[mid] && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

//func search(nums []int, target int) int {
//
//	low := 0
//	high := len(nums) - 1
//	for low <= high {
//		mid := (low + high) / 2
//		if nums[mid] == target {
//			return mid
//		}
//		if nums[low] <= target && target <= nums[mid] {
//			high = mid - 1
//		} else if target < nums[high] && target > nums[mid] {
//			low = mid + 1
//		} else if nums[low] < nums[mid] && target < nums[low] {
//			low = mid + 1
//		} else if nums[high] > nums[mid] && target > nums[high] {
//			low = mid - 1
//		}
//	}
//
//	return -1
//}
func nextPermutation(nums []int) {
	length := len(nums) - 1
	breakout := false
	breakOutIndex := 0
	for i := length; i > 0; i-- {
		if nums[i-1] < nums[i] {
			breakOutIndex = i - 1
			breakout = true
			break
		}
	}

	for j := length; j >= breakOutIndex+1; j-- {
		if nums[breakOutIndex] < nums[j] {
			nums[breakOutIndex], nums[j] = nums[j], nums[breakOutIndex]
			break
		}
	}
	index := breakOutIndex + 1

	if !breakout {
		index = 0
	}
	for j := len(nums) - 1; index < j; index, j = index+1, j-1 {
		nums[index], nums[j] = nums[j], nums[index]
	}
	fmt.Println(nums)
}

func combinationSum(candidates []int, target int) [][]int {
	out := [][]int{}

	if len(candidates) == 0 {
		return out
	}
	sum(&out, make([]int, 0), candidates, 0, target)
	return out

}

func sum(out *[][]int, temp, nums []int, start, target int) {

	if target < 0 || start > len(nums) {
		return
	}

	if target == 0 {
		cpyTmp := make([]int, len(temp))
		copy(cpyTmp, temp)
		*out = append(*out, cpyTmp)
	}
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		sum(out, temp, nums, i, target-nums[i])
		temp = temp[:len(temp)-1]
	}
}

func Permutation(nums []int) [][]int {
	out := make([][]int, 0)
	freq := make([]bool, len(nums))
	doPermutation(&out, make([]int, 0), nums, freq)
	return out
}

func doPermutation(out *[][]int, temps, nums []int, freq []bool) {
	if len(temps) == len(nums) {
		copyPermutations := make([]int, len(temps))
		copy(copyPermutations, temps)
		*out = append(*out, copyPermutations)
	}

	for i, val := range nums {

		if !freq[i] {
			freq[i] = true
			temps = append(temps, val)
			doPermutation(out, temps, nums, freq)
			temps = temps[:len(temps)-1]
			freq[i] = false
		}

	}
}

func groupAnagrams1(strs []string) [][]string {
	out := make([][]string, 0)
	freq := make([]bool, len(strs))
	doGroupAnagrams(&out, make([]string, 0), strs, freq)
	return out
}

func doGroupAnagrams(out *[][]string, temps, nums []string, freq []bool) {
	if len(temps) == len(nums) {
		copyPermutations := make([]string, len(temps))
		copy(copyPermutations, temps)
		*out = append(*out, copyPermutations)
	}

	for i, val := range nums {

		if !freq[i] {
			freq[i] = true
			temps = append(temps, val)
			doGroupAnagrams(out, temps, nums, freq)
			temps = temps[:len(temps)-1]
			freq[i] = false
		}

	}
}

func twoSum(nums []int, target int) (out []int) {
	mapTest := make(map[int]int)
	out = make([]int, 2)
	for j, v := range nums {

		if i, ok := mapTest[target-v]; ok {
			out = []int{i, j}
			return out
		} else {
			mapTest[v] = j
		}
	}
	return out
}

func lengthOfLongestSubstring(s string) int {
	mapTemp := make(map[rune]int)
	startIndex := 0
	maxLength := 0
	for lIndex, lValue := range s {
		if mValue, ok := mapTemp[lValue]; ok {
			startIndex = max(startIndex, mValue+1)
		}
		mapTemp[lValue] = lIndex
		maxLength = max(maxLength, lIndex+1-startIndex)
	}
	return 0 //int(math.Min())
}
func myPow(x float64, n int) float64 {
	pow := float64(1)
	minus := false
	if n < 0 {
		minus = true
		n = -n
	}
	for n > 0 {
		if (n & 1) != 0 {
			pow = pow * x
		}
		x = x * x
		n >>= 1
	}
	if minus {
		pow = 1 / pow
	}
	return pow
	//pow := float64(1)
	//for n > 0 {
	//	if n&1 != 0 {
	//		pow = pow * x
	//	}
	//	x = x * x
	//	n >>= 1
	//}
	//return pow
}

//func topKFrequent(nums []int, k int) []int {
//	a := make(map[int]int)
//
//	for _, v := range nums {
//		a[v]++
//	}
//
//	biarr := make([][]int, len(nums)+1)
//	for k, v := range a {
//		biarr[v] = append(biarr[v], k)
//	}
//
//	out := make([]int, 0, k)
//	for i := len(biarr) - 1; i >= 0; i-- {
//		for _, ans := range biarr[i] {
//			if k > 0 {
//				out = append(out, ans)
//				k--
//			}
//		}
//	}
//	return out
//}

func firstUniqChar(s string) int {
	arr := make([]int, 26)
	for _, ch := range s {
		arr[ch-'a']++
	}
	for i, ch := range s {
		if arr[ch-'a'] == 1 {
			return i
		}
	}

	return -1

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}

func uniquePathsWithObstacles(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return 0
	}
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	//start containing obstacle is handled earlier
	dp[0][0] = 1
	//filling base case for first row and column in dp array
	for j := 1; j < cols; j++ {
		if grid[0][j] == 0 && dp[0][j-1] == 1 {
			//current grid and last grid is empty in first column
			//then there exists a path
			dp[0][j] = 1
		} else {
			//otherwise no path due to obstacles in first column
			dp[0][j] = 0
		}
	}
	for i := 1; i < rows; i++ {
		if grid[i][0] == 0 && dp[i-1][0] == 1 {
			//current grid and last grid is empty in first row
			//then there exists a path
			dp[i][0] = 1
		} else {
			//otherwise no path due to obstacles in first row
			dp[i][0] = 0
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if grid[i][j] == 0 {
				//if particular grid has no obstacle
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[rows-1][cols-1]
}

func minAddToMakeValid(s string) int {
	ma := make(map[rune]rune)

	ma[']'] = ')'
	i := 1
	for i >= 0 && i < len(s) {
		if s[i] == ')' && i >= 1 {
			s = s[0:i-1] + s[i+1:]
			i--
		} else {
			i++
		}

	}

	return len(s)
}

