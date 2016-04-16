package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
	"github.com/antoine-richard/geoquiz"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		question := geoquiz.NextQuestion()
		displayQuestion(question)

		fmt.Println("\nType the numbers one by one")
		var playerAnswers []geoquiz.Answer

		for i := 0; i < question.NumberOfCorrectAnswers; i++ {
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			text = strings.TrimSpace(text)

			if strings.EqualFold(text, "exit") {
				fmt.Println("\nSee ya!")
				os.Exit(1)
			} else {
				answerIndex, err := strconv.Atoi(text)
				if err == nil {
					playerAnswers = append(playerAnswers, question.Answers[answerIndex])
				}
			}
		}

		result := geoquiz.AnswerCurrentQuestion(playerAnswers)
		if result {
			fmt.Println("Correct answer :)\n")
		} else {
			fmt.Println("Bad answer :(\n")
			displayCorrectAnswers(question)
		}

		fmt.Println("Press enter for next question, type exit to quit")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.EqualFold(strings.TrimSpace(text), "exit") {
			fmt.Println("\nSee ya!")
			os.Exit(1)
		}

	}
}

func displayQuestion(question geoquiz.Question) {
	fmt.Println()
	fmt.Println(question.Statement, ":\n")
	for index, answer := range question.Answers {
		fmt.Println(index, "-", answer.CountryName)
	}
}

func displayCorrectAnswers(question geoquiz.Question) {
	fmt.Println("Correct answers are:")
	for index, answer := range question.Answers {
		if answer.Correct {
			fmt.Println(index, "-", answer.CountryName)
		}
	}
}