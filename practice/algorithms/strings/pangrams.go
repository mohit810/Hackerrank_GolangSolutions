package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func pangrams(s string) string {
	tempDictionary := make(map[string]int) //this is made for only comparison and count
	str := strings.ToLower(s)
	for _, val := range str {
		if string(val) != " " {
			_, newFlag := tempDictionary[string(val)]
			if !newFlag {
				tempDictionary[string(val)] = 0
			}
		}
	}
	if len(tempDictionary) == 26 {
		return "pangram"
	} else {
		return "not pangram"
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

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
