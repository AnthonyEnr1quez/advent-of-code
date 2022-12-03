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

	var commonPrioritySum, groupPrioritySum, groupCounter int
	var groupItems []string
	for scanner.Scan() {
		items := scanner.Text()

		groupCounter++
		groupItems = append(groupItems, items)
		if groupCounter == 3 {
			var commonGroupItem rune
			for _, v := range groupItems[0] {
				val := string(v)
				if strings.Contains(groupItems[1], val) && strings.Contains(groupItems[2], val) {
					commonGroupItem = v
					break
				}
			}

			if unicode.IsUpper(commonGroupItem) {
				for i, v := range alphabetUpper {
					if commonGroupItem == v {
						groupPrioritySum += i + 27
					}
				}
			} else {
				for i, v := range alphabetLower {
					if commonGroupItem == v {
						groupPrioritySum += i + 1
					}
				}
			}

			groupCounter = 0
			groupItems = nil
		}

		half := len(items) / 2
		compartment1 := items[:half]
		compartment2 := items[half:]

		var commonItem rune
		for _, v := range compartment1 {
			value := string(v)
			if strings.Contains(compartment2, value) {
				commonItem = v
				break
			}
		}

		if unicode.IsUpper(commonItem) {
			for i, v := range alphabetUpper {
				if commonItem == v {
					commonPrioritySum += i + 27
				}
			}
		} else {
			for i, v := range alphabetLower {
				if commonItem == v {
					commonPrioritySum += i + 1
				}
			}
		}
	}

	fmt.Println(commonPrioritySum)
	fmt.Println(groupPrioritySum)
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
