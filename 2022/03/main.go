package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var alphabetUpper, alphabetLower = alphabets()

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var commonPrioritySum, groupPrioritySum int
	var groupItems []string
	for scanner.Scan() {
		items := scanner.Text()

		half := len(items) / 2
		compartment1 := items[:half]
		compartment2 := items[half:]
		for _, v := range compartment1 {
			if strings.Contains(compartment2, string(v)) {
				commonPrioritySum += getItemPriority(v)
				break
			}
		}

		groupItems = append(groupItems, items)
		if len(groupItems) == 3 {
			for _, v := range groupItems[0] {
				val := string(v)
				if strings.Contains(groupItems[1], val) && strings.Contains(groupItems[2], val) {
					groupPrioritySum += getItemPriority(v)
					break
				}
			}

			groupItems = nil
		}
	}

	fmt.Println("Compartment priority sum:", commonPrioritySum)
	fmt.Println("Group priority sum:", groupPrioritySum)
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

func getItemPriority(item rune) (priority int) {
	if unicode.IsUpper(item) {
		for i, v := range alphabetUpper {
			if item == v {
				priority = i + 27
			}
		}
	} else {
		for i, v := range alphabetLower {
			if item == v {
				priority = i + 1
			}
		}
	}
	return
}
