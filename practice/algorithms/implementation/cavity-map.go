package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var grid [100]string
	for y := 0; y < n; y++ {
		fmt.Scanf("%s", &grid[y])
	}

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			var char uint8
			char = grid[y][x]
			if x == 0 || x == n-1 || y == 0 || y == n-1 {
				fmt.Printf("%c", grid[y][x])
			} else {
				if char > grid[y-1][x] && char > grid[y+1][x] && char > grid[y][x-1] && char > grid[y][x+1] {
					fmt.Printf("X")
				} else {
					fmt.Printf("%c", grid[y][x])
				}
			}
		}
		fmt.Println()
	}
}
