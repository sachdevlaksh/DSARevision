/*
LeetCode Problem #459: Repeated Substring Pattern
Difficulty: Easy

Given a string s, check if it can be constructed by taking a substring of it and adding multiple copies of the substring together.
*/

package LC459


func repeatedSubstringPattern(s string) bool {
    n := len(s)
    for i := 1; i <= n/2; i++ {
        if n%i != 0 {
            continue
        }
        substring := s[:i]

        k := 0
        subFullString := ""
        for k <= len(s) {
            subFullString = repeatSubstring(subFullString, substring)
            k = len(subFullString)
            if subFullString == s {
                return true
            }
        }
    }

    return false

}

func repeatSubstring(s, sub string) string {
    return s + sub
}

func RepeatedSubstringPattern(s string) bool {
    n := len(s)
    if n <= 1{
        return false
    }
    ss := s + s

    ss = ss[1:2*n-1]
    if Containers(ss,s){
        return true
    }

    return false
}


func Containers(s, sub string) bool {
    n, m := len(s), len(sub)

    if m == 0 {
        return true // empty substring always exists
    }
    if m > n {
        return false // longer substring can't exist in shorter string
    }

    for i := 0; i <= n-m; i++ {
        match := true
        for j := 0; j < m; j++ {
            if s[i+j] != sub[j] {
                match = false
                break
            }
        }
        if match {
            return true
        }
    }

    return false
}