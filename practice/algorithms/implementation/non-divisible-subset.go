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

func nonDivisibleSubset(k int32, s []int32) int32 {
	n := len(s)
	f := make([]int32, k)
	//count of the remainders with value 0<=remainder<k
	for i := 0; i < n; i++ {
		f[s[i]%k]++
	}

	// if k % 2 == 0 {f[k/2] = int32(math.Min(float64(f[k/2]), float64(1)))}

	res := int32(math.Min(float64(f[0]), 1))

	for i := 1; i <= int(k/2); i++ {
		if i != int(k)-i {
			res += int32(math.Max(float64(f[i]), float64(f[int(k)-i])))
		}
	}
	if k%2 == 0 {
		res++
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

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
