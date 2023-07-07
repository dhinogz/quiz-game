package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type QuizGame struct {
	Index    int
	Question string
	Answer   string
}

func createQuizGame(data [][]string) []QuizGame {
	var quizList []QuizGame
	for i, line := range data {
		var quiz QuizGame
		quiz.Index = i + 1
		for j, field := range line {
			if j == 0 {
				quiz.Question = strings.TrimSpace(field)
			} else if j == 1 {
				quiz.Answer = strings.TrimSpace(field)
			}
		}
		quizList = append(quizList, quiz)
	}
	return quizList
}

func main() {
	fmt.Println("Hello quiz game!")

	// TODO: Parse csv_file flag to read CSV file (add a default csv file)
	fileName := "problems.csv"
	// TODO: Parse time_limit flag to limit quiz time (add a default quiz time limit) Goroutines are needed here

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	quizList := createQuizGame(data)

	var score int
	reader := bufio.NewReader(os.Stdin)
	for _, q := range quizList {
		fmt.Printf("Problem #%d: %s = ", q.Index, q.Question)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			log.Fatal(err)
		}
		if input == "" {
			fmt.Println("No answer")
		}
		if input == q.Answer {
			score++
		}
	}
	fmt.Println(score)
}
