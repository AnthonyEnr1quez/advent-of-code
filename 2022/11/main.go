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

	monkeys := make(map[int]*Monkey)
	inputs := []string{}

	for scanner.Scan() {
		input := scanner.Text()

		if input != "" {
			inputs = append(inputs, input)
		} else {
			name, monkey := newMonkey(inputs)
			monkeys[name] = monkey

			// reset inputs
			inputs = []string{}
		}
	}

	// add the last monkey, not added during scanning
	name, monkey := newMonkey(inputs)
	monkeys[name] = monkey

	for round := 1; round <= 20; round++ {
		for name := 0; name < len(monkeys); name++ {
			monkey = monkeys[name]

			for _, item := range monkey.Items {
				var worryLevel, opValue, destination int

				if monkey.Operation.Value == "old" {
					opValue = item
				} else {
					val, _ := strconv.Atoi(monkey.Operation.Value)
					opValue = val
				}

				if monkey.Operation.Symbol == "*" {
					worryLevel = item * opValue
				} else {
					worryLevel = item + opValue
				}

				monkey.Inspections++
				postInspectionWorryLevel := worryLevel / 3

				if postInspectionWorryLevel%monkey.Test.Value == 0 {
					destination = monkey.Test.TrueDest
				} else {
					destination = monkey.Test.FalseDest
				}

				monkeys[destination].Items = append(monkeys[destination].Items, postInspectionWorryLevel)
			}
			monkey.Items = []int{}
		}
	}

	inspections := make([]int, len(monkeys))
	for i := range inspections {
		inspections[i] = monkeys[i].Inspections
	}
	sort.Ints(inspections)

	monkeyBusiness := inspections[len(inspections)-1] * inspections[len(inspections)-2]

	fmt.Println(monkeyBusiness)
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
