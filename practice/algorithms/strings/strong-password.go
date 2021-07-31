package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func minimumNumber(n int, password string) int {
	count := 0
	var matchCap = regexp.MustCompile("([A-Z]+)")
	var matchLowerCap = regexp.MustCompile("([a-z]+)")
	var matchNums = regexp.MustCompile("([0-9]+)")
	var matchSymbols = regexp.MustCompile("([!@#$%^&*()\\-+]+)")
	if caps := matchCap.MatchString(password); !caps {
		count++
	}
	if lowerCaps := matchLowerCap.MatchString(password); !lowerCaps {
		count++
	}
	if nums := matchNums.MatchString(password); !nums {
		count++
	}
	if symbols := matchSymbols.MatchString(password); !symbols {
		count++
	}

	if n < 6 {
		if 6-n > count {
			count = 6 - n
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	password := readLine(reader)

	answer := minimumNumber(n, password)

	fmt.Fprintf(writer, "%d\n", answer)

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
