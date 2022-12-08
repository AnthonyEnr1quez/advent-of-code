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

	var grid [][]string
	counter := 0
	for scanner.Scan() {
		input := scanner.Text()

		if grid == nil {
			len := len(input)
			grid = make([][]string, len)
			for i := range grid {
				grid[i] = []string{}
			}
		}
		trees := strings.Split(input, "")
		grid[counter] = append(grid[counter], trees...)
		counter++
	}

	directions := []func(int, int, int, int, [][]string, bool) (bool, int){up, down, left, right}
	visibleTrees := 0
	maxScenicScore := 0
	for y := range grid {
		length := len(grid[y])
		switch y {
		case 0:
			visibleTrees += length
		case length - 1:
			visibleTrees += length
		default:
			visibleTrees += 2 // edge trees
			for x := 1; x < length-1; x++ {
				scenicScore := 1
				visible := false
				for _, direction := range directions {
					visibility, distance := from(direction, x, y, grid)
					if !visible {
						visible = visibility
					}

					scenicScore *= distance
				}

				if visible {
					visibleTrees++
				}

				if scenicScore > maxScenicScore {
					maxScenicScore = scenicScore
				}
			}
		}
	}

	fmt.Println("Visible trees:", visibleTrees)
	fmt.Println("Max scenic score:", maxScenicScore)
}

func from(direction func(int, int, int, int, [][]string, bool) (bool, int), x, y int, grid [][]string) (bool, int) {
	height, _ := strconv.Atoi(grid[y][x])
	return direction(x, y, height, 0, grid, true)
}

func up(x, y, height, distance int, grid [][]string, visible bool) (bool, int) {
	for i := y - 1; i >= 0; i-- {
		treeHeight, _ := strconv.Atoi(grid[i][x])
		if visible {
			distance++
			if height <= treeHeight {
				visible = false
			}
		}
	}
	return visible, distance
}

func down(x, y, height, distance int, grid [][]string, visible bool) (bool, int) {
	for i := y + 1; i < len(grid); i++ {
		treeHeight, _ := strconv.Atoi(grid[i][x])
		if visible {
			distance++
			if height <= treeHeight {
				visible = false
			}
		}
	}
	return visible, distance
}

func right(x, y, height, distance int, grid [][]string, visible bool) (bool, int) {
	for i := x + 1; i < len(grid[y]); i++ {
		treeHeight, _ := strconv.Atoi(grid[y][i])
		if visible {
			distance++
			if height <= treeHeight {
				visible = false
			}
		}
	}
	return visible, distance
}

func left(x, y, height, distance int, grid [][]string, visible bool) (bool, int) {
	for i := x - 1; i >= 0; i-- {
		treeHeight, _ := strconv.Atoi(grid[y][i])
		if visible {
			distance++
			if height <= treeHeight {
				visible = false
			}
		}
	}
	return visible, distance
}
