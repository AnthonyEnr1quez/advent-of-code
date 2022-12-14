package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/9
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

	dosKnots := simulateRope(2, motions)
	diezKnots := simulateRope(10, motions)

	fmt.Println("Positions tail visited using rope with 2 knots:", dosKnots)
	fmt.Println("Positions tail visited using rope with 10 knots:", diezKnots)
}

type Motion struct {
	Direction string
	Steps     int
}

type Point struct {
	X, Y int
}

func (p *Point) touches(other Point) bool {
	return math.Abs(float64(p.X-other.X)) <= 1 && math.Abs(float64(p.Y-other.Y)) <= 1
}

func (p *Point) moveTowards(other Point) {
	sign := func(focus, other int) (val int) {
		if math.Signbit(float64(other - focus)) {
			val = -1
		} else if other-focus != 0 {
			val = 1
		}
		return
	}

	p.X += sign(p.X, other.X)
	p.Y += sign(p.Y, other.Y)
}

func simulateRope(knots int, motions []Motion) int {
	rope := make([]Point, knots)
	tailVisited := make(map[Point]struct{})

	for _, motion := range motions {
		for i := 1; i <= motion.Steps; i++ {
			switch motion.Direction {
			case "U":
				rope[0].Y++
			case "D":
				rope[0].Y--
			case "R":
				rope[0].X++
			case "L":
				rope[0].X--
			}

			for j := 0; j < knots-1; j++ {
				head := &rope[j]
				tail := &rope[j+1]

				if !head.touches(*tail) {
					tail.moveTowards(*head)
				}
			}

			tailVisited[rope[knots-1]] = struct{}{}
		}
	}

	return len(tailVisited)
}
