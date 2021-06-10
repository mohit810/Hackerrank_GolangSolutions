package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkStatus(n int32) {
	if n%2 == 0 {
		if n > 20 {
			fmt.Print("Not Weird")
		} else if 6 <= n && n <= 20 {
			fmt.Print("Weird")
		} else if 2 <= n && n < 5 {
			fmt.Print("Not Weird")
		}
	} else {
		fmt.Print("Weird")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	checkStatus(N)
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
