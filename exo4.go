package main

import (
	"fmt"
)

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // résultat : 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))    // résultat : 0
}
