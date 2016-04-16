package geoquiz

type Question struct {
	Statement string
	NumberOfCorrectAnswers int
	CountryCode string
	CountryName string
	Answers []Answer
}

type Answer struct {
	CountryCode string
	CountryName string
	Correct bool
}

var numberOfCorrectAnswers = 2
var totalNumberOfAnswers = 3 * numberOfCorrectAnswers

var score = 0
var currentQuestion Question

func NextQuestion() Question {
	currentQuestion = getBorderingCountriesQuestion(numberOfCorrectAnswers, totalNumberOfAnswers)
	return currentQuestion
}

func AnswerCurrentQuestion(answers []Answer) (result bool) {
	result = isAnswerCorrect(answers)
	if result {
		score += 1;
	}
	return
}

func isAnswerCorrect(answers []Answer) bool {
	if (len(answers) == currentQuestion.NumberOfCorrectAnswers) {
		for _, answer := range answers {
			if !answer.Correct {
				return false
			}
		}
		return true
	}
	return false
}


