package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	type Data struct {
		i [][]int32
	}
	ref_matrix := []Data{
		{[][]int32{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}}},
		{[][]int32{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}}},
		{[][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}}},
		{[][]int32{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}}},
		{[][]int32{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}}},
		{[][]int32{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}}},
		{[][]int32{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}}},
		{[][]int32{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}}},
	}
	costList := list.New()
	for _, innerMatrix := range ref_matrix {
		cost := float64(0)
		for j, v := range s {
			for k := range v {
				if innerMatrix.i[j][k] != s[j][k] {
					cost += math.Abs(float64(s[j][k] - innerMatrix.i[j][k]))
				}
			}
		}
		costList.PushBack(cost)
	}
	min := costList.Front().Value.(float64)
	for e := costList.Front(); e != nil; e = e.Next() {
		fmt.Print(" e ", e.Value.(float64))
		if min > e.Value.(float64) {
			min = e.Value.(float64)
		}
	}
	return int32(min)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
