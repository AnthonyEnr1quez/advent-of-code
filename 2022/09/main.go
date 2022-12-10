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

	xCord := 600
	yCord := 600
	var grid Grid
	grid = make([][]*Point, yCord)
	for y := range grid {
		grid[y] = make([]*Point, xCord)
		for x := 0; x < xCord; x++ {
			grid[y][x] = &Point{
				X: x,
				Y: y,
			}
		}
	}

	head := grid[300][300]
	tail := head

	for _, motion := range motions {
		for i := 1; i <= motion.Steps; i++ {
			tail.Visited = true

			switch motion.Direction {
			case "U":
				// Tail below Head
				if tail.Y < head.Y {
					tail = head
					tail.Visited = true
				}
				head = grid[head.Y+1][head.X]
			case "D":
				// Tail above Head
				if tail.Y > head.Y {
					tail = head
					tail.Visited = true
				}
				head = grid[head.Y-1][head.X]
			case "R":
				// Tail left of Head
				if tail.X < head.X {
					tail = head
					tail.Visited = true
				}
				head = grid[head.Y][head.X+1]
			case "L":
				// Tail right of Head
				if tail.X > head.X {
					tail = head
					tail.Visited = true
				}
				head = grid[head.Y][head.X-1]
			}
		}
	}

	fmt.Println(grid.Visited(xCord, yCord))
}

type Grid [][]*Point

type Point struct {
	X, Y    int
	Visited bool
}

type Motion struct {
	Direction string
	Steps     int
}

func (g Grid) Print(xCord, yCord int, head, tail *Point) {
	for y := yCord - 1; y >= 0; y-- {
		for x := 0; x < xCord; x++ {
			if g[y][x] == head {
				fmt.Print("H")
			} else if g[y][x] == tail {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g Grid) Visited(xCord, yCord int) (visited int) {
	for y := yCord - 1; y >= 0; y-- {
		for x := 0; x < xCord; x++ {
			if g[y][x].Visited {
				visited++
			}
		}
	}
	return
}
