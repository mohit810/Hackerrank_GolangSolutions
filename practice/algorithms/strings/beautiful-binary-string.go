package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
100
0100101010100010110100100110110100011100111110101001011001110111110000101011011111011001111100011101


10
*/

func beautifulBinaryString(b string, n int32) (count int) {

	if n%3 == 0 {
		for i := 0; i < len(b); i++ {
			if i != len(b)-2 {
				if string(b[i]) == "0" && string(b[i+1]) == "1" && string(b[i+2]) == "0" {
					count++
				}
			} else {
				break
			}
		}
	} else {
		if n%3 == 1 {
			for i := 0; i < len(b)-1; i += 3 {
				if string(b[i]) == "0" && string(b[i+1]) == "1" && string(b[i+2]) == "0" {
					count++
				}
			}
			if string(b[len(b)-3]) == "0" && string(b[len(b)-2]) == "1" && string(b[len(b)-1]) == "0" {
				count++
			}
		} else if n%3 == 2 {

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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	b := readLine(reader)

	result := beautifulBinaryString(b, n)

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
