package LC844

func backspaceCompare(s string, t string) bool {

	st1 := stack{}
	st2 := stack{}

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			st1.pop()
		} else {
			st1.push(s[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			st2.pop()
		} else {
			st2.push(t[i])
		}
	}

	if len(st1) != len(st2) {
		return false
	} else {
		for i := 0; i < len(st1); i++ {
			if st1[i] != st2[i] {
				return false
			}
		}
	}

	return true
}

type stack []byte

func (s *stack) pop() {
	if len(*s) == 0 {
		return
	} else {
		*s = (*s)[:len(*s)-1]
	}
}
func (s *stack) push(x byte) {
	*s = append(*s, x)
}
