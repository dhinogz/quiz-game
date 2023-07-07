package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// TODO: Parse csv_file flag to read CSV file (add a default csv file)
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	flag.Parse()
	// TODO: Parse time_limit flag to limit quiz time (add a default quiz time limit) Goroutines are needed here

	f, err := os.Open(*csvFilename)
	if err != nil {
		exit("Failed to read CSV file.")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		exit("Failed to parse CSV file.")
	}
	problems := createProblems(lines)

	score := 0
	for i, q := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, q.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == q.answer {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d\n", score, len(problems))
}

type problem struct {
	question string
	answer   string
}

func createProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(msg string) {
	log.Fatal(msg)
	os.Exit(1)
}
