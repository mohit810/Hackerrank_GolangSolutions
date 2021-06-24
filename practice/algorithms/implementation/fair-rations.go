package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func fairRations(B []int32) string {
	posibility := true
	count := 0
	for i := 0; i < len(B); i++ {
		if B[i]%2 == 0 {
			continue
		} else {
			if i+1 != len(B) {
				B[i] += 1
				B[i+1] += 1
				count += 2
			} else if i == len(B)-1 && B[i-1]%2 != 0 {
				B[i] += 1
				B[i-1] += 1
				count += 2
			} else {
				posibility = false
			}
		}
	}
	if posibility {
		return strconv.Itoa(count)
	} else {
		return "NO"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)

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
