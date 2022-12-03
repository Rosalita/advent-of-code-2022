package main

import (
	"bufio"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"io"
	"log"
	"os"
	"strings"
)

func getInput() string {
	file, err := os.Open("../../input/3.txt")
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

	var items []rune
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		mid := len(line) / 2

		first_rucksack_string := line[:mid]
		second_rucksack_string := line[mid:]

		first_rucksack := []rune(first_rucksack_string)
		second_rucksack := []rune(second_rucksack_string)

		intersection := intersect.Simple(first_rucksack, second_rucksack)
		rune := intersection[0].(rune)
		items = append(items, rune)
	}
	total := 0
	for _, item := range items {
		priority := calcPriority(item)
		total += priority
	}

	fmt.Println(total)
}

func calcPriority(r rune) int {
	if r > 96 {
		// lowercase bytes are 97 - 122
		// to transform these values to 1 - 26, substract 96
		return int(r - 96)
	} else {
		// uppercase bytes are 65 - 90
		// to transform these values to 27 - 52, subtract 38.
		return int(r - 38)
	}
}
