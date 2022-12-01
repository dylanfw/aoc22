package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e) // skip error handling
	}
}

func readElfCounts(filename string) []int {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var elves []int
	var currentElf int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, currentElf)
			currentElf = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		check(err)
		currentElf += calories
	}
	elves = append(elves, currentElf) // last block has no trailing blank line

	return elves
}

func sum(ints []int) int {
	var total int
	for _, n := range ints {
		total += n
	}
	return total
}

func main() {
	elves := readElfCounts("input.txt")
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Println("(Part 1) Most Calories Held by an Elf:", elves[0])
	fmt.Println("(Part 2) Calories Held by Top 3 Elves:", sum(elves[0:3]))
}
