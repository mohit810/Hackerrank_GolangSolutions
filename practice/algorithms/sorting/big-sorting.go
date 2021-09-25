package sorting

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func MergeSort(slice []big.Int) []big.Int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func merge(left, right []big.Int) []big.Int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]big.Int, size, size)

	for k := 0; k < size; k++ {

		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if comparator := left[i].Cmp(&right[j]); comparator == -1 {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func bigSorting(unsorted []big.Int) []big.Int {
	// Write your code here
	result := MergeSort(unsorted)
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var unsorted []big.Int

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		n, _ := new(big.Int).SetString(unsortedItem, 10)
		unsorted = append(unsorted, *n)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem.String())

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
