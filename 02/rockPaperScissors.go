package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shape string

const (
	Rock     = "Rock"
	Paper    = "Paper"
	Scissors = "Scissors"
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
		"X": Rock,     // Used for Part 1 Only
		"Y": Paper,    // Used for Part 1 Only
		"Z": Scissors, // Used for Part 1 Only
	}[shapeCode]
}

func translateOutcomeCode(opponentShape Shape, outcomeCode string) Shape {
	return map[string]map[Shape]Shape{
		"X": {Rock: Scissors, Paper: Rock, Scissors: Paper},
		"Y": {Rock: Rock, Paper: Paper, Scissors: Scissors},
		"Z": {Rock: Paper, Paper: Scissors, Scissors: Rock},
	}[outcomeCode][opponentShape]
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

		opponentShape := translateShapeCode(shapeCodes[0])
		// selfShape := translateShapeCode(shapeCodes[1]) 				// Part 1
		selfShape := translateOutcomeCode(opponentShape, shapeCodes[1]) // Part 2

		rounds = append(rounds, Round{
			opponentShape: opponentShape,
			selfShape:     selfShape,
		})
	}

	return rounds
}

func main() {
	rounds := readStrategyGuide("input.txt")
	fmt.Println("Total Game Score:", scoreGame(rounds))
}
