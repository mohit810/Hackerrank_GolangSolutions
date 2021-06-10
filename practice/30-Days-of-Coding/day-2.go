package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func round(num float64) float64 {
	if num < 0 {
		return math.Ceil(num - 0.5)
	}
	return math.Floor(num + 0.5)
}

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	tipCost := (meal_cost * float64(tip_percent)) / 100
	totalTax := (meal_cost * float64(tax_percent)) / 100
	totalCost := meal_cost + tipCost + totalTax
	fmt.Print(int(round(totalCost)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	meal_cost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
