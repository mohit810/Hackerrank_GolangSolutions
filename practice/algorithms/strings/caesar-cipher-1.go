package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const alphabets string = "abcdefghijklmnopqrstuvwxyz"

var dictionary map[string]string

func rotate(s []rune, k int) {
	dictionary = make(map[string]string)
	cipher := string(append(s[k:], s[0:k]...))
	for i := 0; i < len(alphabets); i++ {
		dictionary[string(alphabets[i])] = string(cipher[i])
	}
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func caesarCipher(s string, k int) string {
	var matchAlphabets = regexp.MustCompile("([A-Za-z])")
	rotate([]rune(alphabets), k%26)
	var encryptMsg strings.Builder
	for i := 0; i < len(s); i++ {
		if matchAlphabets.MatchString(string(s[i])) {
			if upperCase := IsUpper(string(s[i])); upperCase {
				character := dictionary[strings.ToLower(string(s[i]))]
				encryptMsg.WriteString(strings.ToUpper(character))
			} else {
				encryptMsg.WriteString(dictionary[string(s[i])])
			}
		} else {
			encryptMsg.WriteString(string(s[i]))
		}
	}
	return encryptMsg.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	_ = int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int(kTemp)

	result := caesarCipher(s, k)

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
