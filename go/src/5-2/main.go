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
	file, err := os.Open("../../input/5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return string(data)
}

type stack struct {
	slice []string
}

func (s *stack) len() int {
	return len(s.slice)
}

func (s *stack) push(str string) {
	s.slice = append(s.slice, str)
}

func (s *stack) insert(index int, str string) {
	if len(s.slice) == index {
		s.slice = append(s.slice, str)
		return
	}
	s.slice = append(s.slice[:index+1], s.slice[index:]...)
	s.slice[index] = str
}

func (s *stack) pop() string {
	last := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return last
}

func main() {
	i := getInput()
	scanner := bufio.NewScanner(strings.NewReader(i))

	var contentLines []string
	var moveLines []string
	var numStacks int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "[") {
			contentLines = append(contentLines, line)
			continue
		}
		if strings.Contains(line, "move") {
			moveLines = append(moveLines, line)
			continue
		}
		line = strings.ReplaceAll(line, " ", "")
		numStacks = len(line)
	}

	var stacks []stack

	for i := 0; i < numStacks; i++ {
		stack := stack{}
		stacks = append(stacks, stack)
	}

	contentLines = reverse(contentLines)

	for _, line := range contentLines {
		for i := 0; i < numStacks; i++ {
			index := 1 + (i * 4)
			content := string(line[index])
			if content == " " {
				continue
			}
			stacks[i].push(content)
		}
	}

	for _, line := range moveLines {
		moveInstruction := strings.Split(line, " ")
		n, err := strconv.Atoi(moveInstruction[1])
		if err != nil {
			log.Fatal(err)
		}
		from, err := strconv.Atoi(moveInstruction[3])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(moveInstruction[5])
		if err != nil {
			log.Fatal(err)
		}
		insert_index := stacks[to-1].len()
		for i := 0; i < n; i++ {
			content := stacks[from-1].pop()
			stacks[to-1].insert(insert_index, content)
		}
	}
	for i := 0; i < numStacks; i++ {
		fmt.Printf("%s", stacks[i].pop())
	}
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
