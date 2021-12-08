package solutions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Day1() (counter1, counter2 int) {
	arr := []int{}
	file, err := os.Open("data/day1_1.txt")
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
		valueToAdd, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, valueToAdd)
	}
	for i := 0; i < len(arr); i++ {
		if i+1 <= len(arr)-1 {
			if arr[i] < arr[i+1] {
				counter1++
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		if i+3 <= len(arr)-1 {
			if arr[i]+arr[i+1]+arr[i+2] < arr[i+1]+arr[i+2]+arr[i+3] {
				counter2++
			}
		}
	}
	return counter1, counter2
}
