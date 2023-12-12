package main

import (
	"fmt"
	"os"
	"strings"
)

func day4() {
	data, err := os.ReadFile("day4Input.txt")
	if err != nil {
		panic("Couldn't read file")
	}

	values := strings.Split(string(data), "\n")
	values = values[:len(values)-1]
	cards := make([]int, len(values))
	for i := range cards {
		cards[i] = 1
	}

	for i, value := range values {
		currentNumMatched := 0
		card := strings.Split(value, ": ")[1]
		numbers := strings.Split(card, " | ")
		current := numbers[1]
		winnings := numbers[0]
		winningsMap := make(map[string]bool)

		for _, winning := range strings.Split(winnings, " ") {
			if winning != "" {
				winningsMap[winning] = true
			}
		}

		for _, token := range strings.Split(current, " ") {
			if _, ok := winningsMap[token]; ok {
				delete(winningsMap, token)
				currentNumMatched++
			}
		}

		currentCardNums := cards[i]

		for j := i + 1; j < currentNumMatched+i+1; j++ {
			cards[j] += 1 * currentCardNums
		}

	}
	sum := 0
	for _, card := range cards {
		sum += card
	}
	fmt.Println("SUM", sum)

}
