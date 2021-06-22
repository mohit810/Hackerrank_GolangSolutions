package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func serviceLane(width []int32, cases [][]int32) []int32 {
	var minWidth []int32
	for j := 0; j < len(cases); j++ {
		min := int32(1000)
		for i := int(cases[j][0]); i <= int(cases[j][1]); i++ {
			if width[i] < min {
				min = width[i]
			}
		}
		minWidth = append(minWidth, min)
	}
	return minWidth
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	tTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	t := int32(tTemp)

	widthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var width []int32

	for i := 0; i < int(n); i++ {
		widthItemTemp, err := strconv.ParseInt(widthTemp[i], 10, 64)
		checkError(err)
		widthItem := int32(widthItemTemp)
		width = append(width, widthItem)
	}

	var cases [][]int32
	for i := 0; i < int(t); i++ {
		casesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var casesRow []int32
		for _, casesRowItem := range casesRowTemp {
			casesItemTemp, err := strconv.ParseInt(casesRowItem, 10, 64)
			checkError(err)
			casesItem := int32(casesItemTemp)
			casesRow = append(casesRow, casesItem)
		}

		if len(casesRow) != 2 {
			panic("Bad input")
		}

		cases = append(cases, casesRow)
	}

	result := serviceLane(width, cases)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
