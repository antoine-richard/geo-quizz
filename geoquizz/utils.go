package geoquizz

import (
	"math/rand"
	"github.com/pariz/gountries"
)

func goodAnswersToArray(countries []gountries.Country) (answers []Answer) {
	for _, country := range countries {
		answers = append(answers, Answer{
			CountryCode: country.Codes.Alpha3,
			CountryName: country.Name.Common,
			Correct: true})
	}
	return
}

func badAnswersMapToArray(countryNamesMap map[string]gountries.Country) (answers []Answer) {
	for _, country := range countryNamesMap {
		answers = append(answers, Answer{
			CountryCode: country.Codes.Alpha3,
			CountryName: country.Name.Common,
			Correct: false})
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
