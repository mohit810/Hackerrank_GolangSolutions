package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func cavityMap(grid []string) []string {
	var newGrid []string
	for i := 0; i < len(grid); i++ {
		item := grid[i]
		firstInt, _ := strconv.Atoi(string(item[0]))
		lastInt, _ := strconv.Atoi(string(item[len(item)-1]))
		var text strings.Builder
		for j := 0; j < len(item); j++ {
			innerItem, _ := strconv.Atoi(string(item[j]))
			if j != 0 || j != len(item)-1 {
				if innerItem > firstInt && innerItem > lastInt {
					leftItem, _ := strconv.Atoi(string(item[j-1]))
					rightItem, _ := strconv.Atoi(string(item[j+1]))
					if innerItem > leftItem && innerItem > rightItem {
						text.WriteString("x")
					} else {
						text.WriteString(string(item[j]))
					}
				} else {
					text.WriteString(string(item[j]))
				}
			}
		}
		newGrid = append(newGrid, text.String())
		text.Reset()
	}
	return newGrid
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

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
