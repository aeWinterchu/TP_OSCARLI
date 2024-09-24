package main

import (
	"fmt"
)

func Ft_max_substring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if index, found := charIndex[s[i]]; found && index >= start {
			start = index + 1
		}
		charIndex[s[i]] = i
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	fmt.Println(Ft_max_substring("abcabcbb")) // resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(Ft_max_substring("bbbbb")) // resultat : 1
	// "b" est la plus grande sous chaine
}
