package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2022/day/6
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var dataStream string
	for scanner.Scan() {
		dataStream = scanner.Text()
	}

	packet := processDataStream(dataStream, 4)
	message := processDataStream(dataStream, 14)

	fmt.Println("start-of-packet marker:", packet)
	fmt.Println("start-of-message marker:", message)
}

func processDataStream(dataStream string, sequenceLength int) (processedChars int) {
	cuzArrays := sequenceLength - 1
	processedChars = cuzArrays
	for i := 0; i < len(dataStream)-cuzArrays; i++ {
		chars := make([]string, sequenceLength)
		for i2 := 0; i2 <= cuzArrays; i2++ {
			chars[i2] = string(dataStream[i+i2])
		}

		processedChars++

		var flag bool
		for i, ch := range chars {
			for i2, chInner := range chars {
				if i != i2 && ch == chInner {
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}

		if !flag {
			break
		}
	}

	return
}
