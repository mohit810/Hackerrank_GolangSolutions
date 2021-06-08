package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countingValleys(steps int32, path string) (count int32) {
	var currentPos int32
	seaLevel := true

	for i := int32(0); i < steps; i++ {
		if path[i] == 'U' {
			currentPos++
		} else {
			currentPos--
		}
		if seaLevel && currentPos == -1 {
			count += 1
		}

		seaLevel = currentPos == 0
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

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
