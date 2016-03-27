package main

import (
	"fmt"
	"github.com/antoine-richard/geo-quizz/geoquizz"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		question := geoquizz.NextQuestion()
		display(question)

		fmt.Println("\nPress enter to see another question, type exit to quit")

		text, _ := reader.ReadString('\n')

		if strings.EqualFold(strings.TrimSpace(text), "exit") {
			fmt.Println("\nSee ya!")
			break
		}
	}
}

func display(question geoquizz.Question) {
	fmt.Print(question.Statement)
	fmt.Println(":\n")

	for _, answer := range question.Answers {
		fmt.Println("- ", answer.CountryName)
		if answer.Correct {
			defer fmt.Println("- ", answer.CountryName)
		}
	}

	fmt.Println("\nCorrect answers:")
}