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

func splitNumber(num int64) []int64 {
	str := strconv.Itoa(int(num))
	var center int
	var res []int64
	var firstHalf string
	var secondHalf string
	if len(str) > 0 {
		center = int(math.Floor(float64(len(str) / 2)))
		for i := 0; i < center; i++ {
			firstHalf += string(str[i])
		}
		for i := center; i < len(str); i++ {
			secondHalf += string(str[i])
		}
	}

	if firstHalf == "" {
		firstHalf = "0"
	}
	if secondHalf == "" {
		secondHalf = "0"
	}
	FIRST_HALF, err := strconv.Atoi(firstHalf)
	checkError(err)
	SECOND_HALF, err := strconv.Atoi(secondHalf)
	checkError(err)
	res = append(res, int64(FIRST_HALF))
	res = append(res, int64(SECOND_HALF))
	return res
}

func kaprekarNumbers(p int64, q int64) {
	var parentSlice []int32
	for p <= q {
		var kaprekarNumbersSlice []int64
		kaprekarNumbersSlice = splitNumber(p * p)
		if kaprekarNumbersSlice[0]+kaprekarNumbersSlice[1] == p {
			fmt.Printf("%d ", kaprekarNumbersSlice[0]+kaprekarNumbersSlice[1])
			parentSlice = append(parentSlice, int32(p))
		}
		p += 1
	}
	if len(parentSlice) == 0 {
		fmt.Printf("INVALID RANGE")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	kaprekarNumbers(pTemp, qTemp)
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
