package sorting

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func insertionSort2(n int, intSlice *[]int32) {

	for i := 1; i < n; i++ {
		key := (*intSlice)[i]
		j := i - 1

		for j > -1 && (*intSlice)[j] > key {
			(*intSlice)[j+1] = (*intSlice)[j]
			j -= 1
		}
		(*intSlice)[j+1] = key
		Printer(intSlice, &n)
	}
}

func Printer(intSlice *[]int32, n *int) {
	for i := 0; i <= *n-1; i++ {
		if i != *n-1 {
			fmt.Printf("%d ", (*intSlice)[i])
		} else {
			fmt.Printf("%d\n", (*intSlice)[i])
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < n; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	insertionSort2(n, &arr)
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
