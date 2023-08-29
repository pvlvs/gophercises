package main

import (
	"strings"
	"unicode"
)

func camelcase(s string) int32 {
	var count int32 = 0
	for _, v := range s {
		char := string(v)
		if char == strings.ToUpper(char) {
			count++
		}
	}

	return count + 1
}

func caesarCipher(s string, k int32) string {
	ret := ""

	for _, v := range s {
		if v+k > 90 && v <= 90 {
			n := v + k - 90
			v = 64 + n 
		} else if v+k > 122 && v <= 122 {
			n := v + k - 122
			v = 94 + n + k
		} else if unicode.IsLetter(v) {
			v += k
		}

		ret += string(v)
	}

	return ret
}
