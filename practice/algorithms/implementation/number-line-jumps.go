package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) (value string) {
	n := int32(0)
	for n < 10000 {
		if x1+v1*n == x2+v2*n {
			value = "YES"
			break
		} else if n == 9999 {
			value = "NO"
			break
		}
		n++
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	x1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	x1 := int32(x1Temp)

	v1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	v1 := int32(v1Temp)

	x2Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	x2 := int32(x2Temp)

	v2Temp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	v2 := int32(v2Temp)

	result := kangaroo(x1, v1, x2, v2)

	fmt.Fprintf(writer, "%s\n", result)

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
