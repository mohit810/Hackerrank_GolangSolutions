package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func flatlandSpaceStations(n int32, c []int) int {
	sort.Ints(c)
	first, last := c[0], int(n)-c[len(c)-1]-1
	d := max(first, last)
	for i := 0; i < len(c)-1; i++ {
		d = max((c[i+1]-c[i])/2, d)
	}
	return d
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int

	for i := 0; i < int(m); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int(cItemTemp)
		c = append(c, cItem)
	}

	result := flatlandSpaceStations(n, c)

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
