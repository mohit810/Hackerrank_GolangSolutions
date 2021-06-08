package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func migratoryBirds(arr []int32) (bird int32) {
	max := int32(0)
	val := make(map[int32]int32)
	var maxCounts []int32
	for i := 0; i < len(arr); i++ {
		_, ok := val[arr[i]]
		if !ok {
			val[arr[i]] = 1
		} else {
			val[arr[i]] += 1
		}

		if val[arr[i]] > max {
			max = val[arr[i]]
		}
	}

	for j := 0; j < len(arr); j++ {
		if val[arr[j]] == max {
			maxCounts = append(maxCounts, arr[j])
		}
	}

	bird = maxCounts[0]
	for i := 1; i < len(maxCounts); i++ {
		if bird > maxCounts[i] {
			bird = maxCounts[i]
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

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
