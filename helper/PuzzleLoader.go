package helper

import (
	"bufio"
	"log"
	"os"
)

func LoadPuzzle(puzzleFilepath string) ([]byte, error) {
	log.Printf("Loadging file: %s", puzzleFilepath)

	content, err := os.ReadFile(puzzleFilepath)
	if err != nil {
		return nil, err
	}
	return content, nil

}

func LoadPuzzleLineByLine(puzzleFilepath string) ([]string, error) {

	var slice []string

	reader, err := os.Open(puzzleFilepath)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return slice, err
}
