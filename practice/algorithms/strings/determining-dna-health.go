package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SingleGene struct {
	Position, Value int
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

	genesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var genes []string

	for i := 0; i < int(n); i++ {
		genesItem := genesTemp[i]
		genes = append(genes, genesItem)
	}

	healthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var health []int

	for i := 0; i < int(n); i++ {
		healthItemTemp, err := strconv.ParseInt(healthTemp[i], 10, 64)
		checkError(err)
		healthItem := int(healthItemTemp)
		health = append(health, healthItem)
	}

	positions := make(map[string][]int)
	values := make(map[string][]SingleGene)

	for i := 0; i < len(genes); i++ {
		gene := genes[i]
		values[gene] = append(values[gene], SingleGene{i, health[i]})
		for j := 0; j < len(gene); j++ {
			subGene := gene[:j+1]
			positions[subGene] = append(positions[subGene], i)
		}
	}

	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s := int32(sTemp)

	minVal, maxVal := -1, -1
	for sItr := 0; sItr < int(s); sItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		first := int(firstTemp)

		lastTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		last := int(lastTemp)

		d := firstMultipleInput[2]

		total := 0
		for i := 0; i < len(d); i++ {
			for j := i + 1; j <= len(d); j++ {
				sub := d[i:j]
				if subPositions, ok := positions[sub]; ok {
					possible := false
					for _, p := range subPositions {
						if first > p {
							continue
						}
						if last >= p {
							possible = true
							break
						} else {
							possible = false
							break
						}
					}
					if !possible {
						break
					}
				} else {
					break
				}
				for _, value := range values[sub] {
					if first > value.Position {
						continue
					}
					if last >= value.Position {
						total += value.Value
					} else {
						break
					}
				}
			}
		}

		if minVal == -1 || total < minVal {
			minVal = total
		}
		if maxVal == -1 || total > maxVal {
			maxVal = total
		}
	}
	fmt.Fprintf(writer, "%d %d", minVal, maxVal)
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
