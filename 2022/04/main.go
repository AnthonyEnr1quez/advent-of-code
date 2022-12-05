package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/4
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fullyContained int
	for scanner.Scan() {
		val := scanner.Text()
		pairs := strings.Split(val, ",")

		elf1 := strings.Split(pairs[0], "-")
		elf2 := strings.Split(pairs[1], "-")

		elf1Min, _ := strconv.Atoi(elf1[0])
		elf1Max, _ := strconv.Atoi(elf1[1])
		elf2Min, _ := strconv.Atoi(elf2[0])
		elf2Max, _ := strconv.Atoi(elf2[1])

		if (elf1Min <= elf2Min && elf1Max >= elf2Max) || (elf2Min <= elf1Min && elf2Max >= elf1Max) {
			fullyContained++
		}
	}
	fmt.Println(fullyContained)
}
