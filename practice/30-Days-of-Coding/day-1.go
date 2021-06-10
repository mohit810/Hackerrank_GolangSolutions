package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var (
		secondInt uint64
		secondDub float64
		secondStr string
	)

	fmt.Scan(&secondInt)
	fmt.Scan(&secondDub)
	scanner.Scan()
	secondStr = scanner.Text()
	fmt.Println(i + secondInt)
	fmt.Printf("%.1f\n", d+secondDub)
	fmt.Printf("%s%s", s, secondStr)
}
