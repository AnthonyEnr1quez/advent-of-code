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

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var counter, stackCount int
	var stacks map[int]*Stack[string]

	for scanner.Scan() {
		counter++
		val := scanner.Text()

		// init the stacks
		if counter == 1 {
			stackCount = (len(val) - (len(val) / 4)) / 3
			stacks = make(map[int]*Stack[string], stackCount)

			for i := 1; i <= stackCount; i++ {
				stacks[i] = &Stack[string]{}
			}
		}

		// build the stacks
		crossStackCounter := 1
		if val != "" && val[:2] != " 1" && val[:4] != "move" {
			for i := 1; i <= len(val); i += 4 {
				crate := val[i : i+1]
				if crate != " " {
					stacks[crossStackCounter].Push(val[i : i+1])
				}

				crossStackCounter++
			}
		}

		// reverse the stacks
		if val == "" {
			for i := range stacks {
				correctOrderStack := Stack[string]{}
				popped, err := stacks[i].Pop()
				for err == nil {
					correctOrderStack.Push(popped)
					popped, err = stacks[i].Pop()
				}
				stacks[i] = &correctOrderStack
			}
		}

		// move the crates
		if val != "" && val[:4] == "move" {
			values := strings.Split(val, " ")
			crateCount, _ := strconv.Atoi(values[1])
			from, _ := strconv.Atoi(values[3])
			to, _ := strconv.Atoi(values[5])

			for i := 1; i <= crateCount; i++ {
				crate, _ := stacks[from].Pop()
				if err != nil {
					log.Fatalln("Stack", from, ",", err.Error())
				}
				stacks[to].Push(crate)
			}
		}
	}

	// get the crates on top of each stack
	topCrates := ""
	for i := 1; i <= len(stacks); i++ {
		crate, err := stacks[i].Peek()
		if err != nil {
			log.Fatalln("Stack", i, ",", err.Error())
		}
		topCrates += crate
	}

	fmt.Print(topCrates)
}

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
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

func (s *Stack[T]) Peek() (val T, err error) {
	len := len(s.vals)
	if len == 0 {
		return val, errors.New("Nothing on stack")
	}

	val = s.vals[len-1]
	return
}
