package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EVEN = iota
	ODD
)

func print(ODD_EVEN int, words string) {
	for i := ODD_EVEN; i < len(words); i += 2 {
		fmt.Printf("%c", words[i])
	}
	fmt.Print(" ")
}

func main() {
	var T int

	fmt.Scan(&T)
	words := make([]string, T)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < T; i++ {
		words[i], _ = reader.ReadString('\n')
		words[i] = strings.TrimSuffix(words[i], "\n")
	}

	for k := range words {
		print(EVEN, words[k])
		print(ODD, words[k])
		fmt.Printf("\n")
	}

}
