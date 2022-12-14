package main

import (
	"bufio"
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

	var motions []Motion
	for scanner.Scan() {
		motion := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(motion[1])

		motions = append(motions, Motion{Direction: motion[0], Steps: steps})
	}

	head := Point{X: 0, Y: 0}
	tail := Point{X: 0, Y: 0}

	visited := make(map[Point]struct{})

	for _, motion := range motions {
		for i := 1; i <= motion.Steps; i++ {
			switch motion.Direction {
			case "U":
				// Tail below Head
				if tail.Y < head.Y {
					tail = head
				}
				head.Y++
			case "D":
				// Tail above Head
				if tail.Y > head.Y {
					tail = head
				}
				head.Y--
			case "R":
				// Tail left of Head
				if tail.X < head.X {
					tail = head
				}
				head.X++
			case "L":
				// Tail right of Head
				if tail.X > head.X {
					tail = head
				}
				head.X--
			}

			visited[tail] = struct{}{}
		}
	}

	fmt.Println(len(visited))
}

type Point struct {
	X, Y int
}

type Motion struct {
	Direction string
	Steps     int
}
