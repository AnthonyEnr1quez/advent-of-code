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

	visibleTrees := 0
	for y := range grid {
		len := len(grid[y])
		switch y {
		case 0:
			visibleTrees += len
		case len - 1:
			visibleTrees += len
		default:
			visibleTrees += 2 // edge trees
			for x := 1; x < len-1; x++ {
				if visibleFrom(up, x, y, grid) || visibleFrom(right, x, y, grid) || visibleFrom(down, x, y, grid) || visibleFrom(left, x, y, grid) {
					visibleTrees++
				}
			}
		}
	}

	fmt.Println(visibleTrees)
}

func visibleFrom(f func(int, int, int, [][]string) bool, x, y int, grid [][]string) bool {
	height, _ := strconv.Atoi(grid[y][x])
	return f(x, y, height, grid)
}

func up(x, y, height int, grid [][]string) bool {
	for i := y - 1; i >= 0; i-- {
		treeHeight, _ := strconv.Atoi(grid[i][x])
		if height <= treeHeight {
			return false
		}
	}
	return true
}

func down(x, y, height int, grid [][]string) bool {
	for i := y + 1; i < len(grid); i++ {
		treeHeight, _ := strconv.Atoi(grid[i][x])
		if height <= treeHeight {
			return false
		}
	}
	return true
}

func right(x, y, height int, grid [][]string) bool {
	for i := x + 1; i < len(grid[y]); i++ {
		treeHeight, _ := strconv.Atoi(grid[y][i])
		if height <= treeHeight {
			return false
		}
	}
	return true
}

func left(x, y, height int, grid [][]string) bool {
	for i := x - 1; i >= 0; i-- {
		treeHeight, _ := strconv.Atoi(grid[y][i])
		if height <= treeHeight {
			return false
		}
	}
	return true
}
