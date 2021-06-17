package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var personName, phoneNumber string

	phoneBook := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &personName, &phoneNumber)
		phoneBook[personName] = phoneNumber
	}

	for scanner.Scan() {
		name := scanner.Text()
		value, answer := phoneBook[name]
		if answer == true {
			fmt.Printf("%s=%s\n", name, value)
		} else {
			fmt.Printf("Not found\n")
		}
	}
}
