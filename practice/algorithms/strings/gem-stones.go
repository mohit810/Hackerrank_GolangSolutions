package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func gemstones(arr []string, n int) int32 {

	gemstonesCount := int32(0)
	dictionary := make(map[string]int)

	for _, value := range arr {
		tempDictionary := make(map[string]int)
		for _, innerValue := range value {
			if _, tempFlag := tempDictionary[string(innerValue)]; !tempFlag {
				tempDictionary[string(innerValue)] += 1
				dictionary[string(innerValue)] += 1
			}
		}
	}

	for _, value := range dictionary {
		if value == n {
			gemstonesCount++
		}
	}

	return gemstonesCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	var arr []string

	for i := 0; i < n; i++ {
		arrItem := readLine(reader)
		arr = append(arr, arrItem)
	}

	result := gemstones(arr, n)

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
