package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/11
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	coolMonkeys := make(map[int]*Monkey)
	scaryMonkeys := make(map[int]*Monkey)
	inputs := []string{}

	for scanner.Scan() {
		input := scanner.Text()

		if input != "" {
			inputs = append(inputs, input)
		} else {
			name, monkey := newMonkey(inputs)
			coolMonkeys[name] = monkey

			name, monkey = newMonkey(inputs)
			scaryMonkeys[name] = monkey

			// reset inputs
			inputs = []string{}
		}
	}

	// add the last coolMonkey, not added during scanning
	name, coolMonkey := newMonkey(inputs)
	coolMonkeys[name] = coolMonkey

	name, scaryMonkey := newMonkey(inputs)
	scaryMonkeys[name] = scaryMonkey

	for round := 1; round <= 10000; round++ {
		for name := 0; name < len(coolMonkeys); name++ {
			coolMonkey = coolMonkeys[name]
			scaryMonkey = scaryMonkeys[name]

			if round <= 20 {
				for _, item := range coolMonkey.Items {
					var worryLevel, opValue, destination int

					if coolMonkey.Operation.Value == "old" {
						opValue = item
					} else {
						val, _ := strconv.Atoi(coolMonkey.Operation.Value)
						opValue = val
					}

					if coolMonkey.Operation.Symbol == "*" {
						worryLevel = item * opValue
					} else {
						worryLevel = item + opValue
					}

					coolMonkey.Inspections++
					postInspectionWorryLevel := worryLevel / 3

					if postInspectionWorryLevel%coolMonkey.Test.Value == 0 {
						destination = coolMonkey.Test.TrueDest
					} else {
						destination = coolMonkey.Test.FalseDest
					}

					coolMonkeys[destination].Items = append(coolMonkeys[destination].Items, postInspectionWorryLevel)
				}
				coolMonkey.Items = []int{}
			}

			for _, item := range scaryMonkey.Items {
				var worryLevel, opValue, destination int

				if scaryMonkey.Operation.Value == "old" {
					opValue = item
				} else {
					val, _ := strconv.Atoi(scaryMonkey.Operation.Value)
					opValue = val
				}

				if scaryMonkey.Operation.Symbol == "*" {
					worryLevel = item * opValue
				} else {
					worryLevel = item + opValue
				}

				scaryMonkey.Inspections++
				// postInspectionWorryLevel := worryLevel / 3

				if worryLevel%scaryMonkey.Test.Value == 0 {
					destination = scaryMonkey.Test.TrueDest
				} else {
					destination = scaryMonkey.Test.FalseDest
				}

				scaryMonkeys[destination].Items = append(scaryMonkeys[destination].Items, worryLevel)
			}
			scaryMonkey.Items = []int{}
		}
	}

	coolInspections := make([]int, len(coolMonkeys))
	for i := range coolInspections {
		coolInspections[i] = coolMonkeys[i].Inspections
	}
	sort.Ints(coolInspections)

	coolMonkeyBusiness := coolInspections[len(coolInspections)-1] * coolInspections[len(coolInspections)-2]

	fmt.Println(coolMonkeyBusiness)

	scaryInspections := make([]int, len(scaryMonkeys))
	for i := range scaryInspections {
		scaryInspections[i] = scaryMonkeys[i].Inspections
	}
	sort.Ints(scaryInspections)

	scaryMonkeyBusiness := scaryInspections[len(scaryInspections)-1] * scaryInspections[len(scaryInspections)-2]

	fmt.Println(scaryMonkeyBusiness)
}

type Monkey struct {
	Items       []int
	Operation   Operation
	Test        Test
	Inspections int
}

type Operation struct {
	Symbol, Value string
}

type Test struct {
	Value, TrueDest, FalseDest int
}

func newMonkey(inputs []string) (int, *Monkey) {
	name, _ := strconv.Atoi(strings.Split(inputs[0], " ")[1][:1])
	items := parseItems(inputs[1])
	operation := parseOperation(inputs[2])
	test := parseTest(inputs[3], inputs[4], inputs[5])

	return name, &Monkey{
		Items:     items,
		Operation: operation,
		Test:      test,
	}
}

func parseItems(in string) []int {
	strItems := strings.Split(strings.Split(in, ": ")[1], ", ")

	items := make([]int, len(strItems))
	for i, item := range strItems {
		val, _ := strconv.Atoi(item)
		items[i] = val
	}

	return items
}

func parseOperation(in string) Operation {
	ins := strings.Split(strings.Split(in, ": ")[1], " ")

	return Operation{
		Symbol: ins[3],
		Value:  ins[4],
	}
}

func parseTest(test, trueCond, falseCond string) Test {
	value, _ := strconv.Atoi(strings.Split(test, " ")[5])
	trueDest, _ := strconv.Atoi(strings.Split(trueCond, " ")[9])
	falseDest, _ := strconv.Atoi(strings.Split(falseCond, " ")[9])

	return Test{
		Value:     value,
		TrueDest:  trueDest,
		FalseDest: falseDest,
	}
}
