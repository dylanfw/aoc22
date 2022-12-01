package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func check(e error) {
	if e != nil {
		panic(e)		// skip error handling
	}
}

func findMaxCalories(filename string) int {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	maxCalories := 0.0
	currentCalories := 0.0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// new block of calories
			maxCalories = math.Max(maxCalories, currentCalories)
			currentCalories = 0.0
			continue
		}

		calories, err := strconv.ParseFloat(line, 64)
		check(err)
		currentCalories += calories
	}
	maxCalories = math.Max(maxCalories, currentCalories)  // last block is not followed by blank line
	
	return int(maxCalories)
}

func main() {
	fmt.Println(findMaxCalories("input.txt"))
}