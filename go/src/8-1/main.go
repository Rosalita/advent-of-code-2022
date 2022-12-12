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
	visible := 0
	// loop over every tree
	for x, row := range f.trees {
		for y, tree := range row {
			// dont need to check first row as they are all visible
			if x == 0 {
				visible += 1
				continue
			}

			// dont need to check the last row of trees as they are all visible
			if x == len(f.trees)-1 {
				visible += 1
				continue
			}

			// dont need to check first tree in a row as they are all visible
			if y == 0 {
				visible += 1
				continue
			}

			// don't need to check last tree in a row as they are all visivle
			if y == len(f.trees[0])-1 {
				visible += 1
				continue
			}

			if f.isVisible(tree, x, y){
				visible += 1
			}
		}
	}
	fmt.Println(visible)
}

func (f forest) isVisible(tree, x, y int) bool {
	// walk north, x decreases, y stays same
	for i := x - 1; i >= 0; i-- {
		if f.trees[i][y] >= tree {
			break
		}
		if i == 0 {
			return true
		}
	}

	// walk south, x increases, y stays same
	for i := x + 1; i < len(f.trees); i++ {
		if f.trees[i][y] >= tree {
			break
		}
		if i == len(f.trees)-1 {
			return true
		}
	}

	// walk east, y increases, x stays same
	for i := y - 1; i >= 0; i-- {
		if f.trees[x][i] >= tree {
			break
		}
		if i == 0 {
			return true
		}

	}

	// walk west, y decreases, x stays same
	for i := y + 1; i < len(f.trees[0]); i++ {
		if f.trees[x][i] >= tree {
			return false
		}
		if i == len(f.trees[0])-1 {
			return true
		}
	}

	return false
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
