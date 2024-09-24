package main

import (
	"fmt"
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	count := 0
	l := 0
	start := 0
	minLen := math.MaxInt32
	have := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		char := s[r]

		if need[char] > 0 {
			have[char]++
			if have[char] <= need[char] {
				count++
			}
		}

		for count == len(t) {
			if r-l+1 < minLen {
				minLen = r - l + 1
				start = l
			}

			lChar := s[l]
			if need[lChar] > 0 {
				have[lChar]--
				if have[lChar] < need[lChar] {
					count--
				}
			}
			l++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[start : start+minLen]
}

func main() {
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // resultat : ""
}
