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

	score := 0

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")

		oppMove := vals[0]
		myMove := vals[1]

		// A for Rock, B for Paper, and C for Scissors (opp)
		// X for Rock, Y for Paper, and Z for Scissors (me)
		switch myMove {
		case "X":
			score++
			if oppMove == "A" {
				score += 3
			} else if oppMove == "C" {
				score += 6
			}
		case "Y":
			score += 2
			if oppMove == "A" {
				score += 6
			} else if oppMove == "B" {
				score += 3
			}
		case "Z":
			score += 3
			if oppMove == "B" {
				score += 6
			} else if oppMove == "C" {
				score +=3
			}
		}
	}

	fmt.Println(score)
}
