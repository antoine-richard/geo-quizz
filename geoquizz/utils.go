package geoquizz

import (
	"math/rand"
	"github.com/pariz/gountries"
)

func toAnswers(countries []gountries.Country, correct bool) (answers []Answer) {
	for _, country := range countries {
		answers = append(answers, Answer{
			CountryCode: country.Codes.Alpha3,
			CountryName: country.Name.Common,
			Correct: correct})
	}
	return
}

func values(countryMap map[string]gountries.Country) (countryList []gountries.Country) {
	countryList = make([]gountries.Country, 0, len(countryMap))
	for  _, value := range countryMap {
		countryList = append(countryList, value)
	}
	return
}

func shuffle(src []Answer) (dest []Answer) {
	dest = make([]Answer, len(src))
	perm := rand.Perm(len(src))
	for index, value := range perm {
		dest[value] = src[index]
	}
	return
}
