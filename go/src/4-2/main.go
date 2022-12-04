package main

import (
	"bufio"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	file, err := os.Open("../../input/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return string(data)
}

func main() {
	i := getInput()
	scanner := bufio.NewScanner(strings.NewReader(i))

	totalOverlap := 0
	for scanner.Scan() {
		line := scanner.Text()
		str_split := strings.Split(line, ",")

		expanded := expandAssignments(str_split)

		firstElf := expanded[0]
		secondElf := expanded[1]

		intersection := intersect.Simple(firstElf, secondElf)
		if len(intersection) > 0 {
			totalOverlap += 1
		}

	}
	fmt.Println(totalOverlap)
}

func expandAssignments(assignments []string) [][]int {
	var fullAssignments [][]int

	for _, assignment := range assignments {
		secs := strings.Split(assignment, "-")

		from, err := strconv.Atoi(secs[0])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(secs[1])
		if err != nil {
			log.Fatal(err)
		}
		var fullAssignment []int
		for i := from; i <= to; i++ {
			fullAssignment = append(fullAssignment, i)
		}
		fullAssignments = append(fullAssignments, fullAssignment)
	}
	return fullAssignments
}
