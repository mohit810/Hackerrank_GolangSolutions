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

func howManyGames(p int32, d int32, m int32, s int32) (count int32) {
	var cost int32 = 0
	if p > s {
		count = 0
	} else if p == s {
		count = 1
	} else if p < s {
		for (cost + p - (count)*d) < s {
			if cost == 0 {
				cost = p
			} else {
				if p-(count)*d <= m && (cost+m) < s {
					break
				} else if (cost + m) > s {
					break
				} else if p-(count)*d > m {
					cost = cost + p - (count)*d
				}
			}
			count++
		}
		if p-(count)*d <= m && (cost+m) < s {
			count = count + int32(math.Floor(float64(s-cost)/float64(m)))
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	pTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	p := int32(pTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	sTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	s := int32(sTemp)

	answer := howManyGames(p, d, m, s)

	fmt.Fprintf(writer, "%d\n", answer)

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
