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

func appendAndDelete(s string, t string, k int) (possibility string) {
	var i int
	for i = 0; i < int(math.Min(float64(len(s)), float64(len(t)))); i++ {
		if s[i] != t[i] {
			break
		}
	}

	var deletesRequired = len(s) - i
	var addsRequired = len(t) - i
	var minRequired = deletesRequired + addsRequired
	var max = len(s) + len(t)

	if k < minRequired || ((k%2 != minRequired%2) && k < max) {
		possibility = "No"
	} else {
		possibility = "Yes"
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int(kTemp)

	result := appendAndDelete(s, t, k)

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
