package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func weightedUniformStrings(str string, queries []int) []string {
	var returnSlice []string
	maps := make(map[int]int)
	var prevChar byte
	var count int
	for i := 0; i < len(str); i++ {
		if prevChar == str[i] {
			count++
		} else {
			count = 1
			prevChar = str[i]
		}
		if count > maps[int(str[i]-byte('a')+1)] {
			maps[int(str[i]-byte('a')+1)] = count
		}
	}

	for i := 0; i < len(queries); i++ {
		found := false
		for k, v := range maps {
			if queries[i]%k == 0 && queries[i] <= (k*v) {
				found = true
				break
			}
		}
		if found {
			returnSlice = append(returnSlice, "Yes")
		} else {
			returnSlice = append(returnSlice, "No")
		}
	}
	return returnSlice
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

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
