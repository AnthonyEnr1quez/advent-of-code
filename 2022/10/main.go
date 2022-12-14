package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/10
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	registerByCycleValue := make(map[int]int)

	register := 1
	cycle := 0

	rows := [6]string{}
	rowIndex := 0
	colIndex := 0

	for scanner.Scan() {
		instruction := scanner.Text()

		char := "."

		if strings.Contains(instruction, "noop") {
			if colIndex == register || colIndex == register-1 || colIndex == register+1 {
				char = "#"
			}

			rows[rowIndex] += char
			colIndex++

			cycle++
			registerByCycleValue[cycle] = register

			if colIndex > 39 {
				colIndex = 0
				rowIndex++
			}
		} else {
			if colIndex == register || colIndex == register-1 || colIndex == register+1 {
				char = "#"
			}

			rows[rowIndex] += char
			colIndex++

			cycle++
			registerByCycleValue[cycle] = register

			if colIndex > 39 {
				colIndex = 0
				rowIndex++
			}

			value, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

			char := "."
			if colIndex == register || colIndex == register-1 || colIndex == register+1 {
				char = "#"
			}

			rows[rowIndex] += char
			colIndex++

			cycle++
			registerByCycleValue[cycle] = register

			if colIndex > 39 {
				colIndex = 0
				rowIndex++
			}

			register += value
		}
	}

	signalStrengths := 0
	for _, v := range []int{20, 60, 100, 140, 180, 220} {
		signalStrengths += registerByCycleValue[v] * v
	}

	fmt.Println(signalStrengths)

	formattedAnswer := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", rows[0], rows[1], rows[2], rows[3], rows[4], rows[5])
	fmt.Println(formattedAnswer)
}
