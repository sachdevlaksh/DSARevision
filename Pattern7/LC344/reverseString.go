package LC344

func reverseString(s []byte)  {

	n := len(s)
	if n <= 1{
		return
	}
	left := 0
	right := n-1

	for left<= right{
		s[right],s[left] = s[left],s[right]
		left++
		right--
	}
}