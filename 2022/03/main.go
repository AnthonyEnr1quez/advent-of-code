package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	alphabetUpper, alphabetLower := alphabets()

	var prioritySum int
	for scanner.Scan() {
		items := scanner.Text()
		half := len(items) / 2

		compartment1 := items[:half]
		compartment2 := items[half:]

		var common rune
		for _, v := range compartment1 {
			value := string(v)
			if strings.Contains(compartment2, value) {
				common = v
				break
			}
		}
		
		if unicode.IsUpper(common) {
			for i, v := range alphabetUpper {
				if common == v {
					prioritySum += i + 27
				}
			}
		} else {
			for i, v := range alphabetLower {
				if common == v {
					prioritySum += i + 1
				}
			}
		}
	}
	
	fmt.Println(prioritySum)
}

func alphabets() (upper, lower []rune) {
	var ch rune
	for ch = 'A'; ch <= 'Z'; ch++ {
		upper = append(upper, ch)
	}
	for ch = 'a'; ch <= 'z'; ch++ {
		lower = append(lower, ch)
	}
	return
}
