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
		assignments := []*string{&str_split[0], &str_split[1]}

		expandAssignments(assignments)

		if len(*assignments[0]) > len(*assignments[1]) {
			if strings.Contains(*assignments[0], *assignments[1]) {
				totalOverlap += 1
			}
		} else {
			if strings.Contains(*assignments[1], *assignments[0]) {
				totalOverlap += 1
			}
		}
	}
	fmt.Println(totalOverlap)
}

func expandAssignments(assignments []*string) {
	for i, assignment := range assignments {
		secs := strings.Split(*assignment, "-")

		from, err := strconv.Atoi(secs[0])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(secs[1])
		if err != nil {
			log.Fatal(err)
		}

		var sb strings.Builder

		for i := from; i <= to; i++ {
			if i < 10 {
				sb.WriteString("0")
			}
			sb.WriteString(fmt.Sprintf("%d", i))
			sb.WriteString(",")
		}
		full_assignment := sb.String()
		assignments[i] = &full_assignment
	}
}
