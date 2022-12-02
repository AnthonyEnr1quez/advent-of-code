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

		// A for Rock, B for Paper, and C for Scissors (opp)
		// X means you need to lose,
		// Y means you need to end the round in a draw,
		// Z means you need to win.
		// part 2
		switch oppMove {
		case "A":
			if strat == "X" {
				//scissors
				correctStratScore += 3

			} else if strat == "Y" {
				// rock
				correctStratScore += 1
				correctStratScore += 3
			} else if strat == "Z" {
				// paper
				correctStratScore += 2
				correctStratScore += 6
			}
		case "B":
			if strat == "X" {
				// rock
				correctStratScore += 1
			} else if strat == "Y" {
				// paper
				correctStratScore += 2
				correctStratScore += 3
			} else if strat == "Z" {
				// scissors
				correctStratScore += 3
				correctStratScore += 6
			}
		case "C":
			if strat == "X" {
				// paper
				correctStratScore += 2
			} else if strat == "Y" {
				// scissors
				correctStratScore += 3
				correctStratScore += 3
			} else if strat == "Z" {
				// rock
				correctStratScore += 1
				correctStratScore += 6
			}
		}

		// A for Rock, B for Paper, and C for Scissors (opp)
		// X for Rock, Y for Paper, and Z for Scissors (me)
		// part 1
		switch strat {
		case "X":
			incorrectStratScore++
			if oppMove == "A" {
				incorrectStratScore += 3
			} else if oppMove == "C" {
				incorrectStratScore += 6
			}
		case "Y":
			incorrectStratScore += 2
			if oppMove == "A" {
				incorrectStratScore += 6
			} else if oppMove == "B" {
				incorrectStratScore += 3
			}
		case "Z":
			incorrectStratScore += 3
			if oppMove == "B" {
				incorrectStratScore += 6
			} else if oppMove == "C" {
				incorrectStratScore += 3
			}
		}
	}

	fmt.Println("Incorrect Strategy Score:",incorrectStratScore)
	fmt.Println("Correct Strategy Score:",correctStratScore)
}
