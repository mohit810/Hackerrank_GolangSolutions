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

func funnyString(s string) string {

	var differenceSlice []float64
	var reverseDifferenceSlice []float64
	flag := true

	for i := 0; i < len(s); i++ {
		if i != (len(s) - 1) {
			differenceSlice = append(differenceSlice, float64(int(s[i+1])-int(s[i])))
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if i != 0 {
			reverseDifferenceSlice = append(reverseDifferenceSlice, float64(int(s[i])-int(s[i-1])))
		}
	}

	for i, val := range differenceSlice {
		if math.Abs(val) != math.Abs(reverseDifferenceSlice[i]) {
			flag = false
			break
		}
	}

	if flag {
		return "Funny"
	} else {
		return "Not Funny"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := funnyString(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
