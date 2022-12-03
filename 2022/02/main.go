package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Your total score is the sum of your scores for each round.

// The score for a single round is the score for the shape you selected
// (1 for Rock, 2 for Paper, and 3 for Scissors)

// plus the score for the outcome of the round
// (0 if you lost, 3 if the round was a draw, and 6 if you won).

// https://adventofcode.com/2022/day/2
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	incorrectStratScore := 0
	correctStratScore := 0

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")

		oppMove := vals[0]
		strat := vals[1]

		// Correct
		// A for Rock, B for Paper, and C for Scissors (opp)
		// X for lose, Y for draw, and Z for win
		// Incorrect
		// X for Rock, Y for Paper, and Z for Scissors (me)
		switch strat {
		case "X":
			incorrectStratScore++
			switch oppMove {
			case "A":
				//scissors
				incorrectStratScore += 3
				correctStratScore += 3
			case "B":
				// rock
				correctStratScore += 1
			case "C":
				// paper
				incorrectStratScore += 6
				correctStratScore += 2
			}
		case "Y":
			incorrectStratScore += 2
			switch oppMove {
			case "A":
				// rock
				incorrectStratScore += 6

				correctStratScore += 1
				correctStratScore += 3
			case "B":
				// paper
				incorrectStratScore += 3

				correctStratScore += 2
				correctStratScore += 3
			case "C":
				// scissors
				correctStratScore += 3
				correctStratScore += 3
			}
		case "Z":
			incorrectStratScore += 3
			switch oppMove {
			case "A":
				// paper
				correctStratScore += 2
				correctStratScore += 6
			case "B":
				// scissors
				incorrectStratScore += 6

				correctStratScore += 3
				correctStratScore += 6
			case "C":
				// rock
				incorrectStratScore += 3

				correctStratScore += 1
				correctStratScore += 6
			}
		}
	}

	fmt.Println("Incorrect Strategy Score:", incorrectStratScore)
	fmt.Println("Correct Strategy Score:", correctStratScore)
}
