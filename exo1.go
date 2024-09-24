package main

import (
	"fmt"
)

func Ft_coin(coins []int, amount int) int {
	min := amount + 1
	var neededAmount func(int, int)
	neededAmount = func(amount int, count int) {
		if amount == 0 {
			if count < min {
				min = count
			}
			return
		}
		if amount < 0 {
			return
		}
		for _, coin := range coins {
			neededAmount(amount-coin, count+1)
		}
	}

	neededAmount(amount, 0)
	if min == amount+1 {
		return -1
	}
	return min
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
}
