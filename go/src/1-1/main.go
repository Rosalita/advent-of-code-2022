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
	file, err := os.Open("../../input/1.txt")
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

	var caloriesPerElf []int
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += calories
			continue
		}
		caloriesPerElf = append(caloriesPerElf, total)
		total = 0
	}

	caloriesPerElf = append(caloriesPerElf, total)

	var max int
	for i, _ := range caloriesPerElf {
		if i == 0 || caloriesPerElf[i] > max {
			max = caloriesPerElf[i]
		}
	}

	fmt.Println(max)
}
