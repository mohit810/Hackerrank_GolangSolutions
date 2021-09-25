package sorting

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func insertionSort1(intSlice *[]int32) {
	var newElement = (*intSlice)[len(*intSlice)-1]
	n := len(*intSlice)
	for i := n - 1; i >= 0; i-- {
		if i != 0 {
			if (*intSlice)[i-1] > newElement {
				(*intSlice)[i] = (*intSlice)[i-1]
				Printer(intSlice, &n)
			} else {
				(*intSlice)[i] = newElement
				Printer(intSlice, &n)
				break
			}
		} else {
			(*intSlice)[i] = newElement
			Printer(intSlice, &n)
		}
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
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	insertionSort1(&arr)
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
