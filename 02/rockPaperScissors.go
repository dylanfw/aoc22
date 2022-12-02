package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shape string

const (
	Unknown  Shape = "Unknown"
	Rock           = "Rock"
	Paper          = "Paper"
	Scissors       = "Scissors"
)

type Round struct {
	opponentShape Shape
	selfShape     Shape
}

func check(e error) {
	if e != nil {
		panic(e) // skip error handling
	}
}

func translateShapeCode(shapeCode string) Shape {
	return map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}[shapeCode]
}

func scoreRound(round Round) int {
	var shapeScores = map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}

	var roundScores = map[Shape]map[Shape]int{
		// Self vs Opponent
		Rock:     {Rock: 3, Paper: 0, Scissors: 6},
		Paper:    {Rock: 6, Paper: 3, Scissors: 0},
		Scissors: {Rock: 0, Paper: 6, Scissors: 3},
	}

	var score int = 0
	score += shapeScores[round.selfShape]
	score += roundScores[round.selfShape][round.opponentShape]
	return score
}

func scoreGame(rounds []Round) int {
	var score int = 0
	for _, round := range rounds {
		score += scoreRound(round)
	}
	return score
}

func readStrategyGuide(filename string) []Round {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var rounds []Round
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		shapeCodes := strings.Split(line, " ")
		rounds = append(rounds, Round{
			opponentShape: translateShapeCode(shapeCodes[0]),
			selfShape:     translateShapeCode(shapeCodes[1]),
		})
	}

	return rounds
}

func main() {
	rounds := readStrategyGuide("input.txt")
	fmt.Println("(Part 1) Total Game Score:", scoreGame(rounds))
}
