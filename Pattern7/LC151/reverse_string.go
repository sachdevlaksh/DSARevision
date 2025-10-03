package LC151

import (
	"strings"
)

func ReverseWords3(s string) string {
	if len(s) == 0 {
		return s
	}
	charArray := make([]string, 0)
	for string(s[0]) == " " {
		s = s[1:]
	}

	for _, ch := range s {
		if ch == ' ' && charArray[len(charArray)-1] == " " {
			continue
		} else {
			charArray = append(charArray, string(ch))
		}
	}
	charArray = append([]string{" "}, charArray...)
	outwords := make([]string, 0)

	for string(charArray[len(charArray)-1]) == " " {
		charArray = charArray[:len(charArray)-1]
	}
	var word strings.Builder
	finalPos := 0
	intialPos := 0
	for i := len(charArray) - 1; i >= 0; i-- {
		if charArray[i] == " " {
			intialPos = i
			for intialPos < finalPos {
				word.WriteString(charArray[intialPos+1])
				intialPos++
			}
			outwords = append(outwords, word.String())
			finalPos = 0
		} else {
			word.Reset()
			if finalPos == 0 {
				finalPos = i
			}
			continue

		}
	}
	outString := ""
	for i := 0; i < len(outwords)-1; i++ {
		outString += outwords[i] + " "
	}

	return outString + outwords[len(outwords)-1]
}

func ReverseWords(s string) string {

	i, n, res := 0, len(s), ""

	for i < n {
		for i < n && s[i] == ' ' {
			i++
		}
		if i == n {
			break
		}
		j := i

		for j < n && s[j] != ' ' {
			j++
		}

		if len(res) == 0 {
			res = s[i:j]

		} else {
			res = s[i:j] + " " + res
		}
		i = j+1

	}
	return res
}
