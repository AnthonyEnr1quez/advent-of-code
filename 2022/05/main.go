package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/5
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var counter, stackCount int
	var cm9000Stacks, cm9001Stacks map[int]*Stack[string]

	for scanner.Scan() {
		counter++
		val := scanner.Text()

		// init the stacks
		if counter == 1 {
			stackCount = (len(val) - (len(val) / 4)) / 3
			cm9000Stacks = make(map[int]*Stack[string], stackCount)
			cm9001Stacks = make(map[int]*Stack[string], stackCount)

			for i := 1; i <= stackCount; i++ {
				cm9000Stacks[i] = &Stack[string]{}
				cm9001Stacks[i] = &Stack[string]{}
			}
		}

		// build the stacks
		crossStackCounter := 1
		if val != "" && val[:2] != " 1" && val[:4] != "move" {
			for i := 1; i <= len(val); i += 4 {
				crate := val[i : i+1]
				if crate != " " {
					cm9000Stacks[crossStackCounter].Push(val[i : i+1])
				}

				crossStackCounter++
			}
		}

		// reverse the stacks
		if val == "" {
			for i := range cm9000Stacks {
				correctOrderStack := Stack[string]{}
				popped, err := cm9000Stacks[i].Pop()
				for err == nil {
					correctOrderStack.Push(popped)
					popped, err = cm9000Stacks[i].Pop()
				}
				cm9000Stacks[i] = &correctOrderStack
			}

			// copy stacks
			for k := range cm9000Stacks {
				cm9001Stacks[k].vals = append(cm9001Stacks[k].vals, cm9000Stacks[k].vals...)
			}
		}

		if val != "" && val[:4] == "move" {
			values := strings.Split(val, " ")
			crateCount, _ := strconv.Atoi(values[1])
			from, _ := strconv.Atoi(values[3])
			to, _ := strconv.Atoi(values[5])

			// move the crates 1 by 1
			for i := 1; i <= crateCount; i++ {
				crate, err := cm9000Stacks[from].Pop()
				if err != nil {
					log.Fatalln("Stack", from, ",", err.Error())
				}
				cm9000Stacks[to].Push(crate)
			}

			// move the crates as groups
			crates, err := cm9001Stacks[from].PopMultiple(crateCount)
			if err != nil {
				log.Fatalln("Stack", from, ",", err.Error())
			}
			cm9001Stacks[to].PushMultiple(crates)
		}
	}

	cm9000topCrates, _ := getTopCrates(cm9000Stacks)
	cm9001topCrates, _ := getTopCrates(cm9001Stacks)

	fmt.Println("CrateMover 9000 top crates:", cm9000topCrates)
	fmt.Println("CrateMover 9001 top crates:", cm9001topCrates)
}

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) PushMultiple(vals []T) {
	for _, v := range vals {
		s.Push(v)
	}
}

func (s *Stack[T]) Pop() (val T, err error) {
	len := len(s.vals)
	if len == 0 {
		return val, errors.New("Nothing on stack")
	}

	val = s.vals[len-1]
	s.vals = s.vals[:len-1]

	return
}

func (s *Stack[T]) PopMultiple(amount int) (vals []T, err error) {
	len := len(s.vals)
	if len-amount < 0 {
		return nil, errors.New(fmt.Sprintf("Not enough values in stack to remove: %d", amount))
	}

	reverseVals := make([]T, amount)
	for i := 0; i < amount; i++ {
		val, err := s.Pop()
		if err != nil {
			return nil, err
		}
		reverseVals[i] = val
	}

	// dirty reverse lol
	for i := amount - 1; i >= 0; i-- {
		vals = append(vals, reverseVals[i])
	}

	return
}

func (s *Stack[T]) Peek() (val T, err error) {
	len := len(s.vals)
	if len == 0 {
		return val, errors.New("Nothing on stack")
	}

	val = s.vals[len-1]
	return
}

func getTopCrates(stacks map[int]*Stack[string]) (topCrates string, err error) {
	for i := 1; i <= len(stacks); i++ {
		crate, peekErr := stacks[i].Peek()
		if peekErr != nil {
			err = errors.New(fmt.Sprint("Stack", i, ",", peekErr.Error()))
		}
		topCrates += crate
	}
	return
}
