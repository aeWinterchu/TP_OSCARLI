package main

import "fmt"

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	for i := 0; i < len(intervals)-1; i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][1] > intervals[j+1][1] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}

	end := intervals[0][1]
	count := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}

func main() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))         // resultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2
}
