package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// https://adventofcode.com/2022/day/1
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var cals []int

	sum := 0
	for scanner.Scan() {
		strVal := scanner.Text()

		if strVal == "" {
			cals = append(cals, sum)
			sum = 0
			continue
		}

		val, err := strconv.Atoi(strVal)
		if err != nil {
			log.Fatalln(err)
		}

		sum += val
	}

	// add the last sum once scanning is done
	cals = append(cals, sum)

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))

	fmt.Println("Most calories:", cals[0])
	fmt.Println("Top 3 calorie sum:", cals[0] + cals[1] + cals[2])
}
