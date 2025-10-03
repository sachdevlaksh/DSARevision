package LC17

func LetterCombinations(digits string) []string {
	wordList := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	out := ""

	ans := make([]string, 0)
	var dfs = func(*[]string, string, string, []string, int) {}
	dfs = func(ans *[]string, out string, digits string, wordList []string, index int) {
		if len(digits) == 0 {
			return
		}
		if len(out) == len(digits) {
			*ans = append(*ans, out)
			return
		}
		buttonDigit := digits[index] - '0'

		buttonWord := wordList[buttonDigit]

		for _,ch := range buttonWord{
			out = out + string(ch)
			dfs(ans, out, digits, wordList, index+1)
			out = out[:len(out)-1]
		}
	}
	dfs(&ans, out, digits, wordList, 0)
	return ans
}
