package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"strings"
)

// Plan:
// construct a data representation of a forest, this is a grid, [][]int
// the trees at the edges of the forest dont need to be checked, only counted
// for each tree not at the egde, check North, South, East and West
// for each direction start walking comparing trees to origin tree
// if reach the edge, then the tree is visible.
// if encounter an equal or bigger tree, stop walking and check next direction

type forest struct {
	trees [][]int
}

func (f forest) countVisbleTrees() {
	maxScore := 0
	// loop over every tree
	for x, row := range f.trees {
		for y, tree := range row {
			score := f.score(tree, x, y)

			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println(maxScore)
}

func (f forest) score(tree, x, y int) int {
	northScore := 0
	// walk north, x decreases, y stays same
	for i := x - 1; i >= 0; i-- {
		northScore += 1
		if f.trees[i][y] >= tree {
			break
		}
	}

	southScore := 0
	// walk south, x increases, y stays same
	for i := x + 1; i < len(f.trees); i++ {
		southScore += 1
		if f.trees[i][y] >= tree {
			break
		}
	}

	eastScore := 0
	// walk east, y increases, x stays same
	for i := y - 1; i >= 0; i-- {
		eastScore += 1
		if f.trees[x][i] >= tree {
			break
		}
	}

	westScore := 0
	// walk west, y decreases, x stays same
	for i := y + 1; i < len(f.trees[0]); i++ {
		westScore += 1
		if f.trees[x][i] >= tree {
			break
		}
	}

	return northScore * southScore * eastScore * westScore
}

func getInput() string {
	file, err := os.Open("../../input/8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return string(data)
}

func main() {
	i := getInput()
	forest := makeForest(i)
	forest.countVisbleTrees()
}

func makeForest(input string) forest {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var forest forest
	for scanner.Scan() {
		line := scanner.Text()
		var trees []int
		for _, rune := range line {
			height, err := strconv.Atoi(string(rune))
			if err != nil {
				log.Fatal(err)
			}
			trees = append(trees, height)
		}
		forest.trees = append(forest.trees, trees)
	}
	return forest
}
