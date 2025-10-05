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


func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapDigitToLetters := map[byte]string{ '2': "abc", '3':  "def", '4':  "ghi",  '5': "jkl", '6':  "mno",  '7': "pqrs", '8':  "tuv",  '9': "wxyz"}
	ans := make([]string, 0)
	var dfs = func( string, int) {}
	dfs = func( out string, index int) {

		if len(out) == len(digits) {
			ans = append(ans, out)
			return
		}

		buttonWord := mapDigitToLetters[digits[index]]
		for _,ch := range buttonWord{
			dfs( out + string(ch), index+1)
		}
	}
	dfs("", 0)
	return ans
}