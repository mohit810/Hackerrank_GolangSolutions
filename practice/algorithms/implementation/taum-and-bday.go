package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(x, y int64) int64 {
	if y < x {
		return y
	}
	return x
}

func taumBday(b int64, w int64, bc int64, wc int64, z int64) int64 {
	costB := min(bc, wc+z) * b
	costW := min(wc, bc+z) * w
	return costB + costW
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int64(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		bTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		b := int64(bTemp)

		wTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		w := int64(wTemp)

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		bcTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		bc := int64(bcTemp)

		wcTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)
		wc := int64(wcTemp)

		zTemp, err := strconv.ParseInt(secondMultipleInput[2], 10, 64)
		checkError(err)
		z := int64(zTemp)

		result := taumBday(b, w, bc, wc, z)

		fmt.Fprintf(writer, "%d\n", result)
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
