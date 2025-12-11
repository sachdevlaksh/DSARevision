package LC242

func IsAnagram(s string, t string) bool {

	var charS [26]int
	var charT [26]int

	for _, su := range s{
		charS[su-'a']++
	}
	for _, tu := range t{
		charT[tu-'a']++
	}

	for i, v := range charS{
		if v != charT[i]{
			return false
		}
	}
	return true

}