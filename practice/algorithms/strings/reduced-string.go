package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func superReducedString(s string) string {
	var newString strings.Builder
	deletionFlag := false

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if s[i] == s[i+1] {
				fmt.Println("s[i] ", string(s[i]), " s[i+1] ", string(s[i+1]))
				temp := delChar([]rune(s), i+1)
				newString.WriteString(string(delChar(temp, i)))
				deletionFlag = true
				break
			}
		}
	}

	if s == "" {
		return "Empty String"
	}

	if deletionFlag {
		return superReducedString(newString.String())
	} else {
		return s
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

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
