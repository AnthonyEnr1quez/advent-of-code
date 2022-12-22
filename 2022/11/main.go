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

	goodMonkeys := make(Barrel)
	evilMonkeys := make(Barrel)

	inputs := []string{}

	for scanner.Scan() {
		input := scanner.Text()

		if input != "" {
			inputs = append(inputs, input)
		} else {
			name, monkey := newMonkey(inputs)
			goodMonkeys[name] = monkey

			name, monkey = newMonkey(inputs)
			evilMonkeys[name] = monkey

			// reset inputs
			inputs = []string{}
		}
	}

	// add the last monkey, not added during scanning
	name, goodMonkey := newMonkey(inputs)
	goodMonkeys[name] = goodMonkey

	name, evilMonkey := newMonkey(inputs)
	evilMonkeys[name] = evilMonkey

	commonDenom := 1
	for _, monkey := range evilMonkeys {
		commonDenom *= monkey.Test.Value
	}

	for round := 1; round <= 10000; round++ {
		for name := 0; name < len(goodMonkeys); name++ {
			if round <= 20 {
				goodMonkey = goodMonkeys[name]

				for _, item := range goodMonkey.Items {
					destination, postInspectionWorryLevel := goodMonkey.inspect(item, func(wl int) int { return wl / 3 })

					goodMonkeys[destination].Items = append(goodMonkeys[destination].Items, postInspectionWorryLevel)
				}
				goodMonkey.Items = []int{}
			}

			evilMonkey = evilMonkeys[name]
			for _, item := range evilMonkey.Items {
				destination, postInspectionWorryLevel := evilMonkey.inspect(item, func(wl int) int { return wl % commonDenom })

				evilMonkeys[destination].Items = append(evilMonkeys[destination].Items, postInspectionWorryLevel)
			}
			evilMonkey.Items = []int{}
		}
	}

	goodMonkeyBusiness := goodMonkeys.monkeyBusiness()

	fmt.Println("Level of monkey business after 20 rounds of stuff-slinging simian shenanigans:", goodMonkeyBusiness)

	evilMonkeyBusiness := evilMonkeys.monkeyBusiness()

	fmt.Println("Level of monkey business after 10000 rounds of stuff-slinging simian shenanigans:", evilMonkeyBusiness)
}

type Barrel map[int]*Monkey

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

func (barrel Barrel) monkeyBusiness() int {
	inspections := make([]int, len(barrel))
	for i := range inspections {
		inspections[i] = barrel[i].Inspections
	}
	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
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

func (monkey *Monkey) inspect(item int, manage func(int) int) (destination, postInspectionWorryLevel int) {
	var worryLevel, opValue int

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
	postInspectionWorryLevel = manage(worryLevel)

	if postInspectionWorryLevel%monkey.Test.Value == 0 {
		destination = monkey.Test.TrueDest
	} else {
		destination = monkey.Test.FalseDest
	}

	return
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
