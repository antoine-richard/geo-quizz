package geoquizz

import (
	"github.com/pariz/gountries"
	"fmt"
	"math/rand"
	"time"
	"errors"
	"log"
)

var query = gountries.New()
var countriesByBorders = tidyCountriesByBorders()

type Question struct {
	Statement string
	CountryName string
	Answers []Answer
}

type Answer struct {
	CountryName string
	Correct bool
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func tidyCountriesByBorders() (countriesByBorders map[int][]gountries.Country) {
	countriesByBorders = make(map[int][]gountries.Country)
	for _, country := range query.Countries {
		borders := len(country.Borders)
		countriesByBorders[borders] = append(countriesByBorders[borders], country)
	}
	return
}

func getQuestion(numberOfBorders int, totalNumberOfAnswers int) Question {
	var questionCountry gountries.Country
	var answers []Answer
	var answersError, countryError error

	for ok := true; ok; ok = (answersError != nil) {
		//questionCountry, _ := query.FindCountryByAlpha("NIC")
		questionCountry, countryError = pickACountry(countriesByBorders, numberOfBorders)
		if countryError != nil {
			log.Fatal(countryError)
		}

		answers, answersError = computeAnswers(questionCountry, totalNumberOfAnswers)
		if answersError != nil {
			log.Println(answersError)
		}
	}

	statement := fmt.Sprintf("Pick %v's %v bordering countries", questionCountry.Name.Common, numberOfBorders)
	return Question{statement, questionCountry.Name.Common, answers}
}

// TODO: write a test
func pickACountry(countriesByBorders map[int][]gountries.Country, numberOfBorders int) (gountries.Country, error) {
	countries := countriesByBorders[numberOfBorders]

	if len(countries) == 0 {
		message := fmt.Sprintf("No more country with %v borders", numberOfBorders)
		return gountries.Country{}, errors.New(message)
	}

	index := rand.Intn(len(countries))
	country := countries[index]

	// TODO: extract this action in the game file?
	// removing the picked country from the list of questions
	countriesByBorders[numberOfBorders] = removeCountry(countries, index)

	return country, nil
}

func computeAnswers(questionCountry gountries.Country, totalNumberOfAnswers int) ([]Answer, error) {
	correctBorderingCountries := questionCountry.BorderingCountries()

	badAnswersMap := make(map[string]bool)
	// adding bad answers
	for _, country := range correctBorderingCountries {
		for _, answersNeighbor := range country.BorderingCountries() {
			if (answersNeighbor.Codes.Alpha3 != questionCountry.Codes.Alpha3)  {
				badAnswersMap[answersNeighbor.Name.Common] = false
			}
		}
	}
	// removing correct answers if present
	for _, country := range correctBorderingCountries {
		delete(badAnswersMap, country.Name.Common)
	}

	// to arrays of Answer
	correctAnswers := countriesToAnswers(correctBorderingCountries)
	badAnswers := namesMapToAnswersArray(badAnswersMap)

	if len(correctAnswers) + len(badAnswers) < totalNumberOfAnswers {
		message := fmt.Sprintf("Not enough bad answers for bordering countries of %v (found %v, wanted %v)",
			questionCountry.Name.Common, len(badAnswers), totalNumberOfAnswers - len(correctAnswers))
		return nil, errors.New(message)
	}

	return limitAndShuffleAnswers(correctAnswers, badAnswers, totalNumberOfAnswers), nil
}

func limitAndShuffleAnswers(correctAnswers []Answer, badAnswers []Answer, totalNumberOfAnswers int) (answers []Answer) {
	answers = correctAnswers

	badAnswers = shuffle(badAnswers)

	i := 0
	for len(answers) < totalNumberOfAnswers {
		answers = append(answers, badAnswers[i])
		i++
	}

	return shuffle(answers)
}
