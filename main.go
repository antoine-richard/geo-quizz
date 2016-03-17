package main

import (
	"github.com/pariz/gountries"
	"fmt"
	"math/rand"
	"time"
	"errors"
	"log"
)

var query = gountries.New()
var countriesByBorders = tidyUpCountriesByBorders()

type Answer struct {
	CountryName string
	Correct bool
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func tidyUpCountriesByBorders() (countriesByBorders map[int][]gountries.Country) {
	countriesByBorders = make(map[int][]gountries.Country)
	for _, country := range query.Countries {
		borders := len(country.Borders)
		countriesByBorders[borders] = append(countriesByBorders[borders], country)
	}
	return
}

func main() {
	var questionCountry gountries.Country
	var answers []Answer
	var err error

	var numberOfBorders = 2
	var totalNumberOfAnswers = 3 * numberOfBorders

	for ok := true; ok; ok = (err != nil) {
		//questionCountry, _ := query.FindCountryByAlpha("NIC")
		questionCountry = pickACountry(countriesByBorders, numberOfBorders)

		answers, err = computeAnswers(questionCountry, totalNumberOfAnswers)
		if err != nil {
			log.Println(err)
		}
	}

	display(questionCountry.Name.Common, numberOfBorders, answers)
}

func pickACountry(countriesByBorders map[int][]gountries.Country, numberOfBorders int) (gountries.Country) {
	countries := countriesByBorders[numberOfBorders]
	return countries[rand.Intn(len(countries))]
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


	// TODO: extract
	// to arrays of Answer
	var correctAnswers, badAnswers []Answer
	for _, country := range correctBorderingCountries {
		correctAnswers = append(correctAnswers, Answer{CountryName: country.Name.Common, Correct: true})
	}
	for countryName := range badAnswersMap {
		badAnswers = append(badAnswers, Answer{CountryName: countryName, Correct: false})
	}


	if len(correctAnswers) + len(badAnswers) < totalNumberOfAnswers {
		message := fmt.Sprintf("Not enough bad answers for bordering countries of %v (found %v, wanted %v)",
			questionCountry.Name.Common, len(badAnswers), totalNumberOfAnswers - len(correctAnswers))
		return nil, errors.New(message)
	}

	return limitAndShuffleAnswers(correctAnswers, badAnswers, totalNumberOfAnswers)
}

func limitAndShuffleAnswers(correctAnswers []Answer, badAnswers []Answer, totalNumberOfAnswers int) (answers []Answer, err error) {
	answers = correctAnswers

	badAnswers = shuffle(badAnswers)

	i := 0
	for len(answers) < totalNumberOfAnswers {
		answers = append(answers, badAnswers[i])
		i++
	}

	return shuffle(answers), nil
}

//TODO extract
func shuffle(src []Answer) (dest []Answer) {
	dest = make([]Answer, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return
}

func display(countryName string , numberOfBorders int, answers []Answer) {
	fmt.Println("\nName", countryName, "'s", numberOfBorders, "bordering countries...\n")
	fmt.Println("Possible answers:")

	for _, answer := range answers {
		fmt.Println("- ", answer.CountryName)
		if answer.Correct {
			defer fmt.Println("- ", answer.CountryName)
		}
	}

	fmt.Println("\nCorrect answers:")
}
