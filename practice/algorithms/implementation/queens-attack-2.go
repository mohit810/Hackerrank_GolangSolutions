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

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	left := c_q - 1
	right := n - c_q
	top := n - r_q
	bottom := r_q - 1
	topLeftDiagonal := int32(math.Min(float64(left), float64(top)))
	topRightDiagonal := int32(math.Min(float64(top), float64(right)))
	bottomLeftDiagonal := int32(math.Min(float64(bottom), float64(left)))
	bottomRightDiagonal := int32(math.Min(float64(bottom), float64(right)))
	for i := int32(0); i < k; i++ {
		row := obstacles[i][0]
		col := obstacles[i][1]
		if row == r_q && col < c_q {
			if c_q-col-1 < left {
				left = c_q - col - 1
			}
		} else if row == r_q && col > c_q {
			if col-c_q-1 < right {
				right = col - c_q - 1
			}
		} else if row > r_q && col == c_q {
			if row-r_q-1 < top {
				top = row - r_q - 1
			}
		} else if row < r_q && col == c_q {
			if r_q-row-1 < bottom {
				bottom = r_q - row - 1
			}
		} else if row > r_q && col < c_q {
			if row-r_q == c_q-col {
				if row-r_q-1 < topLeftDiagonal {
					topLeftDiagonal = row - r_q - 1
				}
			}
		} else if row > r_q && col > c_q {
			if row-r_q == col-c_q {
				if row-r_q-1 < topRightDiagonal {
					topRightDiagonal = row - r_q - 1
				}
			}
		} else if row < r_q && col < c_q {
			if r_q-row == c_q-col {
				if r_q-row-1 < bottomLeftDiagonal {
					bottomLeftDiagonal = r_q - row - 1
				}
			}
		} else if row < r_q && col > c_q {
			if r_q-row == col-c_q {
				if r_q-row-1 < bottomRightDiagonal {
					bottomRightDiagonal = r_q - row - 1
				}
			}
		}
	}
	return left + right + top + bottom + topLeftDiagonal + topRightDiagonal + bottomLeftDiagonal + bottomRightDiagonal
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

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != 2 {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
