package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func libraryFine(d1 int, m1 int, y1 int, d2 int, m2 int, y2 int) int {
	if y1 > y2 {
		return 10000
	} else if m1 > m2 && y1 == y2 {
		return (m1 - m2) * 500
	} else if d1 > d2 && y1 == y2 && m1 == m2 {
		return (d1 - d2) * 15
	} else {
		return 0
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	d1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	d1 := int(d1Temp)

	m1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m1 := int(m1Temp)

	y1Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	y1 := int(y1Temp)

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	d2Temp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	d2 := int(d2Temp)

	m2Temp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	m2 := int(m2Temp)

	y2Temp, err := strconv.ParseInt(secondMultipleInput[2], 10, 64)
	checkError(err)
	y2 := int(y2Temp)

	result := libraryFine(d1, m1, y1, d2, m2, y2)

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
