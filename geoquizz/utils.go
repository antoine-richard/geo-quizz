package geoquizz

import (
	"math/rand"
	"github.com/pariz/gountries"
)

func removeCountry(countries []gountries.Country, index int) []gountries.Country {
	lastIndex := len(countries) - 1
	countries[index] = countries[lastIndex]
	return countries[:lastIndex]
}

func countriesToAnswers(countries []gountries.Country) (answers []Answer) {
	for _, country := range countries {
		answers = append(answers, Answer{CountryName: country.Name.Common, Correct: true})
	}
	return
}

func namesMapToAnswersArray(countryNamesMap map[string]bool) (answers []Answer) {
	for countryName := range countryNamesMap {
		answers = append(answers, Answer{CountryName: countryName, Correct: false})
	}
	return
}

func shuffle(src []Answer) (dest []Answer) {
	dest = make([]Answer, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return
}
