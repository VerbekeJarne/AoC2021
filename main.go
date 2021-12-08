package main

import (
	"AoC/solutions"
	"fmt"
)

func main() {
	day1_1, day1_2 := solutions.Day1()
	fmt.Printf("Solution day 1_1: %d\nSolution day 1_2: %d\n", day1_1, day1_2)
	data2_1 := solutions.Day2_1()
	fmt.Println(data2_1)
	data2_2 := solutions.Day2_2()
	fmt.Println(data2_2)
	data3_1 := solutions.Day3_1()
	fmt.Println(data3_1)
}
