package geoquiz

import (
	"testing"
	"github.com/pariz/gountries"
)

var countryListFixture = []gountries.Country {
	gountries.Country{Codes: gountries.Codes{ Alpha3: "aaa"}},
	gountries.Country{Codes: gountries.Codes{ Alpha3: "bbb"}},
	gountries.Country{Codes: gountries.Codes{ Alpha3: "ccc"}},
}

func TestRemoveCountry(t *testing.T) {
	newList := removeCountry(gountries.Country{Codes: gountries.Codes{ Alpha3: "bbb"}}, countryListFixture)
	if len(newList) != 2 {
		t.Error("removeCountry doesn't remove the country!")
	}
}
