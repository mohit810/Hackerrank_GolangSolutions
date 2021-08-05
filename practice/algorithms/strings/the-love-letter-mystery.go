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

const alphabets string = "abcdefghijklmnopqrstuvwxyz"

var dictionary map[string]int

func theLoveLetterMystery(s string) (counter int) {

	for i := 0; i < len(s); i++ {
		if len(s)%2 != 0 {
			if i != len(s)-1-i {
				firstValue, _ := dictionary[string(s[i])]
				secondValue, _ := dictionary[string(s[len(s)-1-i])]
				counter += int(math.Abs(float64(firstValue - secondValue)))
			} else {
				break
			}
		} else {
			if i != (len(s))/2 {
				firstValue, _ := dictionary[string(s[i])]
				secondValue, _ := dictionary[string(s[len(s)-1-i])]
				counter += int(math.Abs(float64(firstValue - secondValue)))
			} else {
				break
			}
		}
	}
	return
}

func init() {
	dictionary = make(map[string]int)
	for i := 0; i < len(alphabets); i++ {
		dictionary[string(alphabets[i])] = i
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

		result := theLoveLetterMystery(s)

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
