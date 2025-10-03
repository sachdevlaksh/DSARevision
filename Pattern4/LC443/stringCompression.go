package LC443

import (
	"fmt"
	"strconv"
)

func Compress(chars []byte) int {
	charMap := make(map[byte]int)
	for _, char := range chars {
		if _, ok := charMap[char]; !ok {
			charMap[char] = 1
		} else {
			charMap[char] += 1
		}
	}
	outLen := 0
	charOut := make([]byte, 0)
	for k, v := range charMap {
		outLen += 1
		charOut = append(charOut, k)
		if v != 1 {
			if v < 10 {
				outLen += v
				charOut = append(charOut, byte(v))
			}else{
				for _, ch := range []byte(strconv.Itoa(v)) {
					charOut = append(charOut, byte(ch))
				}
			}
		}
	}
	chars = append(charOut,chars ...)
	fmt.Println(string(charOut))
	return outLen
}
