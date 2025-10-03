package main

import "fmt"
func minOperations(s string) int {
	m := make(map[byte]int, 0)
	operations := 0
	for _, ch := range s {
		if ch != rune('a') {
			if operations < 26 - int(byte(ch-'a')){
                x := 26 - int(byte(ch-'a'))
               operations  = x
            }
		}
	}

	return operations
}
