package geoquiz


var numberOfBorders = 2
var totalNumberOfAnswers = 3 * numberOfBorders

var score = 0
var currentQuestion Question

func NextQuestion() Question {
	currentQuestion = getQuestion(numberOfBorders, totalNumberOfAnswers)
	return currentQuestion
}
