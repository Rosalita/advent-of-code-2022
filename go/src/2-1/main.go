package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getInput() string {
	file, err := os.Open("../../input/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return string(data)
}

type shape int

const (
	rock     shape = iota // 0
	paper                 // 1
	scissors              // 2
)

type round struct {
	opponent shape
	player   shape
}

func main() {
	i := getInput()
	scanner := bufio.NewScanner(strings.NewReader(i))

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		opponent_input := string(line[0])
		player_input := string(line[2])

		var thisRound round

		switch opponent_input {
		case "A":
			thisRound.opponent = rock
		case "B":
			thisRound.opponent = paper
		case "C":
			thisRound.opponent = scissors
		}

		switch player_input {
		case "X":
			thisRound.player = rock
		case "Y":
			thisRound.player = paper
		case "Z":
			thisRound.player = scissors
		}

		points := scoreRound(thisRound)
		score += points
	}
	fmt.Println(score)
}

func scoreRound(round round) int {
	points := 0
	switch round.player {
	case rock:
		points += 1
	case paper:
		points += 2
	case scissors:
		points += 3
	}

	if round.opponent == round.player {
		points += 3
	}

	switch round.opponent {
	case rock:
		if round.player == paper {
			points += 6
		}
	case paper:
		if round.player == scissors {
			points += 6
		}
	case scissors:
		if round.player == rock {
			points += 6
		}
	}

	return points
}
