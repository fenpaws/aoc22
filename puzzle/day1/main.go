package day1

import (
	"github.com/fenpaws/aoc22/helper"
	"log"
	"sort"
	"strconv"
)

func Run() {
	log.Print("AOC 22 - Day1")
	raw, err := helper.LoadPuzzleLineByLine("puzzle/day1/puzzle.txt")
	if err != nil {
		log.Print(err)
	}
	//log.Printf("Raw content of Puzzle (calories): %s", content)
	perElfCalories := addElfCalories(raw)
	log.Printf("Combines calories per Elf: %v", perElfCalories)
	elfWithTheMostCalories := elfWithTheMostCalories(perElfCalories)
	log.Printf("Max calories: %v", elfWithTheMostCalories)
	mostCalories := threeElfWithTheMostCalories(perElfCalories)
	log.Printf("Max calories with three Elf: %v", mostCalories)
}

func addElfCalories(calories []string) []int {
	elfCalories := 0
	var elfCaloriesSlice []int
	sizeOfCaloriesSlice := len(calories) - 1

	for i, calories := range calories {

		intCalories, _ := strconv.Atoi(calories)
		elfCalories = elfCalories + intCalories

		if calories == "" || i == sizeOfCaloriesSlice {
			elfCaloriesSlice = append(elfCaloriesSlice, elfCalories)
			elfCalories = 0
		}

	}
	return elfCaloriesSlice

}

func elfWithTheMostCalories(calories []int) int {
	tmpCalories := 0
	for _, calories := range calories {
		if tmpCalories < calories {
			tmpCalories = calories
		}
	}
	return tmpCalories
}

func threeElfWithTheMostCalories(calories []int) int {
	sort.Ints(calories)
	return calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
}
