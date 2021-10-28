package owoifyx

import (
	"math/rand"
	"regexp"
)

const ()

func Owoifxy(unowo string) string {
	kaomoji := []string{
		"(*^ω^)",
		"(◕‿◕✿)",
		"(◕ᴥ◕)",
		"ʕ•ᴥ•ʔ",
		"ʕ￫ᴥ￩ʔ",
		"(*^.^*)",
		"owo",
		"OwO",
		"(｡♥‿♥｡)",
		"uwu",
		"UwU",
		"(*￣з￣)",
		">w<",
		"^w^",
		"(つ✧ω✧)つ",
		"(/ =ω=)/",
	}
	outuwu := []struct {
		rege string
		rep  string
	}{
		{"(?:l|r)", "w"},
		{"(?:L|R)", "W"},
		{"n([aeiou])", "ny$1"},
		{"N([aeiou])|N([AEIOU])", "Ny$1"},
		{"ove", "uv"},

		{"nd(\\s)|nd$", "ndo"},
	}

	for _, x := range outuwu {
		regxp := regexp.MustCompile(x.rege)
		unowo = regxp.ReplaceAllString(unowo, x.rep)

	}

	regxp := regexp.MustCompile("!+")

	unowo = regxp.ReplaceAllStringFunc(unowo, func(string) string {
		return kaomoji[rand.Intn(len(kaomoji))]
	})

	return unowo
}
