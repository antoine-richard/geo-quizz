package geoquiz

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
	CountryCode string
	CountryName string
	Answers []Answer
}

type Answer struct {
	CountryCode string
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
		questionCountry, countryError = pickACountry(numberOfBorders)
		if countryError != nil {
			log.Fatal(countryError)
		}

		answers, answersError = computeAnswers(questionCountry, totalNumberOfAnswers)
		if answersError != nil {
			log.Println(answersError)
		}
	}

	statement := fmt.Sprintf("Pick %v's %v bordering countries", questionCountry.Name.Common, numberOfBorders)
	return Question{statement, questionCountry.Codes.Alpha3, questionCountry.Name.Common, answers}
}

// TODO: write a test
func pickACountry(numberOfBorders int) (country gountries.Country, err error) {
	countries := countriesByBorders[numberOfBorders]

	if len(countries) == 0 {
		message := fmt.Sprintf("No more country with %v borders", numberOfBorders)
		return gountries.Country{}, errors.New(message)
	}

	country = countries[rand.Intn(len(countries))]

	removeFromList(country, countriesByBorders[numberOfBorders])

	return country, nil
}

func computeAnswers(questionCountry gountries.Country, totalNumberOfAnswers int) ([]Answer, error) {
	correctBorderingCountries := questionCountry.BorderingCountries()

	badAnswersMap := make(map[string]gountries.Country)
	// adding bad answers
	for _, country := range correctBorderingCountries {
		for _, answersNeighbor := range country.BorderingCountries() {
			if (answersNeighbor.Codes.Alpha3 != questionCountry.Codes.Alpha3)  {
				badAnswersMap[answersNeighbor.Codes.Alpha3] = answersNeighbor
			}
		}
	}
	// removing correct answers if present
	for _, country := range correctBorderingCountries {
		delete(badAnswersMap, country.Codes.Alpha3)
	}

	// converting countries to answers
	correctAnswers := toAnswers(correctBorderingCountries, true)
	badAnswers := toAnswers(values(badAnswersMap), false)

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

func removeFromList(countryToRemove gountries.Country, countries []gountries.Country) {
	i := 0
	for _, country := range countries {
		if country.Codes.Alpha3 != countryToRemove.Codes.Alpha3 {
			countries[i] = country
			i++
		}
	}
	countriesByBorders[numberOfBorders] = countries[:i]
}
