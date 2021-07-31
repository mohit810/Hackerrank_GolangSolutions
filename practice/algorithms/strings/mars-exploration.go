package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func sosSlices(s string) []string {
	slices := []string{}
	for i := 0; i < len(s); i += +3 {
		slices = append(slices, s[i:i+3])
	}
	return slices
}

func equalCheck(value string) int32 {
	var counter int32
	if string(value[0]) != "S" {
		counter++
	}
	if string(value[1]) != "O" {
		counter++
	}
	if string(value[2]) != "S" {
		counter++
	}
	return counter
}

func marsExploration(s string) int32 {
	var errorCharacterCount int32
	messages := sosSlices(s)
	for _, value := range messages {
		errorCharacterCount += equalCheck(value)
	}
	return errorCharacterCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := marsExploration(s)

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
