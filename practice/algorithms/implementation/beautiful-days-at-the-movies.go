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

func reverse(numbers []int32) []int32 {
	newNumbers := make([]int32, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func splitToDigits(n int32) int32 {
	var ret []int32
	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}
	reverse(ret)
	var op int32 = 1
	var res int32 = 0
	for i := len(ret) - 1; i >= 0; i-- {
		res += ret[i] * op
		op *= 10
	}
	return res
}

func beautifulDays(i int32, j int32, k int32) int {
	var remainder float64
	temp := i
	var days []int32
	for temp <= j {
		hep := splitToDigits(temp)
		remainder = math.Abs(float64(temp-hep)) / float64(k)
		if remainder == math.Trunc(remainder) {
			days = append(days, 1)
		}
		temp++
	}
	return len(days)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
