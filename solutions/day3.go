package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var finalGamma, finalEpsilon []int

func Day3_1() int64 {

	for z := 0; z < 12; z++ {
		amount0, amount1 := retAmount01(z)
		if amount0 > amount1 {
			finalGamma = append(finalGamma, 0)
			finalEpsilon = append(finalEpsilon, 1)
		} else {
			finalGamma = append(finalGamma, 1)
			finalEpsilon = append(finalEpsilon, 0)
		}
	}
	finalGammaString := strings.Trim(strings.Replace(fmt.Sprint(finalGamma), " ", "", -1), "[]")
	finalEpsilonString := strings.Trim(strings.Replace(fmt.Sprint(finalEpsilon), " ", "", -1), "[]")
	finalGammaDecimal, err := strconv.ParseInt(finalGammaString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	finalEpsilonDecimal, err := strconv.ParseInt(finalEpsilonString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return finalGammaDecimal * finalEpsilonDecimal
}

func retAmount01(z int) (amount0, amount1 int) {
	file, err := os.Open("data/day3_1.txt")
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
		tempArr := strings.Split(scanner.Text(), "")
		checkI, err := strconv.Atoi(tempArr[z])
		if err != nil {
			log.Fatal(err)
		}
		if checkI == 0 {
			amount0++
		} else {
			amount1++
		}
	}
	return amount0, amount1
}
