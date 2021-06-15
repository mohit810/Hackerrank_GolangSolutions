package main

import "fmt"

func main() {
	var n int
	mapper := make(map[string]int)
	fmt.Scan(&n)
	var name string
	var phoneNumber int
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scan(&phoneNumber)
		mapper[name] = phoneNumber
	}
	var searchItem string
	for {
		_, err := fmt.Scanf("%s", &searchItem)
		if err != nil {
			break
		}
		if value, ok := mapper[searchItem]; ok {
			fmt.Printf("%s=%d\n", searchItem, value)
		} else {
			fmt.Println("Not found")
		}
	}
}
