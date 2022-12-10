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
	file, err := os.Open("../../input/6.txt")
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
	scanner.Scan()

	index := 0
	to_check := []byte{}

	line := []byte(scanner.Text())

	for _, byte := range line {
		to_check = append(to_check, byte)

		if len(to_check) < 14 {
			index++
			continue
		}

		if is_all_unique(to_check) {
			fmt.Println(index + 1)
			return
		}

		to_check = to_check[1:]
		index++
	}

}

func is_all_unique(vals []byte) bool {
	start_len := len(vals)

	unique := make(map[byte]struct{})
	counter := 0
	for _, val := range vals {
		if _, ok := unique[val]; !ok {
			unique[val] = struct{}{}
			counter++
		}
	}
	return start_len == counter
}
