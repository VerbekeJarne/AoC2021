package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2_1() int {

	var position, depth int

	position = 0
	depth = 0

	file, err := os.Open("data/day2_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		lineValue, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case "forward":
			position = position + lineValue
		case "down":
			depth = depth + lineValue
		case "up":
			depth = depth - lineValue
		}
	}
	fmt.Printf("Position: %d\n", position)
	fmt.Printf("Depth: %d\n", depth)
	return position * depth
}

func Day2_2() int {

	var position, depth, aim int

	position = 0
	depth = 0
	aim = 0

	file, err := os.Open("data/day2_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		lineValue, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case "forward":
			position = position + lineValue
			depth = depth + (aim * lineValue)
		case "down":
			aim = aim + lineValue
		case "up":
			aim = aim - lineValue
		}
	}
	fmt.Printf("Position: %d\n", position)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Aim: %d\n", aim)
	return position * depth
}
