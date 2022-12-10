package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	for scanner.Scan() {
		instruction := scanner.Text()

		if strings.Contains(instruction, "noop") {
			cycle++
			registerByCycleValue[cycle] = register
		} else {
			cycle++
			registerByCycleValue[cycle] = register
			value, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

			cycle++
			registerByCycleValue[cycle] = register
			register += value
		}
	}

	signalStrengths := 0
	for _, v := range []int{20, 60, 100, 140, 180, 220} {
		signalStrengths += registerByCycleValue[v] * v
	}

	fmt.Println(signalStrengths)
}
