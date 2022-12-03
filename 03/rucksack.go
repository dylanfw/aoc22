package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Sack struct {
	leftCompartment  []rune
	rightCompartment []rune
}

func check(e error) {
	if e != nil {
		panic(e) // skip error handling
	}
}

func createSack(line string) Sack {
	items := strings.Split(line, "")
	newSack := Sack{}
	for i, item := range items {
		var compartment *[]rune
		if i < len(items)/2 {
			compartment = &newSack.leftCompartment
		} else {
			compartment = &newSack.rightCompartment
		}

		*compartment = append(*compartment, []rune(item)[0])
	}
	return newSack
}

func groupElves(sacks []Sack) [][]Sack {
	var groupSize int = 3
	var elfGroups [][]Sack
	for i := 0; i < len(sacks); i += groupSize {
		elfGroups = append(elfGroups, sacks[i:i+groupSize])
	}
	return elfGroups
}

func findCommonItem(sack Sack) rune {
	for _, iItem := range sack.leftCompartment {
		for _, jItem := range sack.rightCompartment {
			if iItem == jItem {
				return jItem
			}
		}
	}
	return rune(0)
}

func allSackItems(sack Sack) []rune {
	return append(sack.leftCompartment, sack.rightCompartment...)
}

func findCommonGroupItem(sacks []Sack) rune {
	for _, iItem := range allSackItems(sacks[0]) {
		for _, jItem := range allSackItems(sacks[1]) {
			for _, kItem := range allSackItems(sacks[2]) {
				if iItem == jItem && jItem == kItem {
					return kItem
				}
			}
		}
	}
	return rune(0)
}

func itemPriority(item rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	if item >= 96 && item < 123 {
		// a-z
		return int(item - 96)
	} else if item >= 65 && item < 91 {
		// A-Z
		return int(item - 38)
	}
	return 0
}

func readSackContents(filename string) []Sack {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var sacks []Sack
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sacks = append(sacks, createSack(line))
	}

	return sacks
}

func main() {
	sacks := readSackContents("input.txt")

	sumPriority := 0
	for _, sack := range sacks {
		sumPriority += itemPriority(findCommonItem(sack))
	}
	fmt.Println("(Part 1) Sum of Common Item Priorities:", sumPriority)

	elfGroups := groupElves(sacks)
	sumGroupPriority := 0
	for _, group := range elfGroups {
		sumGroupPriority += itemPriority(findCommonGroupItem(group))
	}
	fmt.Println("(Part 2) Sum of Common Group Item Priorities:", sumGroupPriority)
}
