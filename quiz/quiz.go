package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("/home/qwerty/IdeaProjects/golab2/quiz/quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

func ask(s score, question question, channel chan score) {
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		fmt.Println("Correct!")
		s++
	} else {
		fmt.Println("Incorrect :-(")
	}
	channel <- s
}

func askQuestions(reschan chan score) {
	s := score(0)
	qs := questions()
	channel := make(chan score)

	for _, q := range qs {
		go ask(s, q, channel)
		s = <-channel
	}
	reschan <- s
}

func main() {
	reschan := make(chan score)
	go askQuestions(reschan)

	time.Sleep(500 * time.Millisecond)

	s := <-reschan
	fmt.Println("Final score", s)

	os.Exit(0)
}
