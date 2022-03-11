package main

import "fmt"

// Say that you are a traveler on a 2D grid.
// You begin in the top-left corner and your goal is to travel to the bottom-right corner.
// You may only move down or right.

// In how many ways can you travel to the goal on a grid with dimensions m * n?

type Position struct {
	Left int
	Top  int
}

func gridTravellerMemo(m, n int) int {
	return gridTravellerMemoHelper(m, n, map[Position]int{})
}

func gridTravellerMemoHelper(m, n int, cache map[Position]int) int {
	if m <= 0 || n <= 0 {
		return 0
	} else if m == 1 && n == 1 {
		return 1
	}

	v, ok := cache[Position{
		Left: m,
		Top:  n,
	}]
	if ok {
		return v
	}

	res := gridTravellerMemoHelper(m-1, n, cache) + gridTravellerMemoHelper(m, n-1, cache)
	cache[Position{
		Left: m,
		Top:  n,
	}] = res

	return res
}

func gridTravellerTab(m, n int) int {
	items := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		items[i] = make([]int, n+1)
	}

	if m > 0 && n > 0 {
		items[1][1] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			items[i][j] += items[i-1][j] + items[i][j-1]
		}
	}

	return items[m][n]
}

func main() {
	fmt.Println(gridTravellerMemo(1, 1), gridTravellerTab(1, 1))
	fmt.Println(gridTravellerMemo(2, 3), gridTravellerTab(2, 3))
	fmt.Println(gridTravellerMemo(3, 2), gridTravellerTab(3, 2))
	fmt.Println(gridTravellerMemo(3, 3), gridTravellerTab(3, 3))
	fmt.Println(gridTravellerMemo(18, 18), gridTravellerTab(18, 18))
}
