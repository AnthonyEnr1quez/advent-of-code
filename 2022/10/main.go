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

	var yIdx, xIdx, cycle, signalStrengths int
	register := 1
	crt := [6]string{}

	for scanner.Scan() {
		instruction := scanner.Text()

		iterations := 1
		val := 0

		if !strings.Contains(instruction, "noop") {
			iterations = 2
			val, _ = strconv.Atoi(strings.Split(instruction, " ")[1])
		}

		for i := 1; i <= iterations; i++ {
			pixel := "."

			if xIdx == register || xIdx == register-1 || xIdx == register+1 {
				pixel = "#"
			}

			crt[yIdx] += pixel
			xIdx++

			cycle++
			signalStrengths += signal(cycle, register)

			if xIdx > 39 {
				xIdx = 0
				yIdx++
			}
		}

		register += val
	}

	fmt.Println("Sum of interesting signal strengths:", signalStrengths)

	fmt.Println("\nCRT screen image ↓")
	for _, row := range crt {
		fmt.Println(row)
	}
}

func signal(cycle, register int) (strength int) {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		strength = register * cycle
	}
	return
}
