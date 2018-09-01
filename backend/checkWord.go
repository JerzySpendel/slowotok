package main

import (
	"encoding/json"
	"net/http"
	"strings"
)


func Score(word string) (score int) {
	scoreDict := make(map[string]int)
	scoreDict["A"] = 1
	scoreDict["E"] = 1
	scoreDict["I"] = 1
	scoreDict["N"] = 1
	scoreDict["O"] = 1
	scoreDict["R"] = 1
	scoreDict["S"] = 1
	scoreDict["W"] = 1
	scoreDict["Z"] = 1
	scoreDict["D"] = 2
	scoreDict["C"] = 2
	scoreDict["K"] = 2
	scoreDict["L"] = 2
	scoreDict["M"] = 2
	scoreDict["P"] = 2
	scoreDict["T"] = 2
	scoreDict["Y"] = 2
	scoreDict["U"] = 3
	scoreDict["J"] = 3
	scoreDict["H"] = 3
	scoreDict["G"] = 3
	scoreDict["B"] = 3
	scoreDict["F"] = 5

	total := 0

	for i := 0; i < len(word); i++ {
		total += scoreDict[strings.ToUpper(string(word[i]))]
	}

	return total
}

func CheckWord(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	word := req.Form.Get("word")

	if len(word) == 0 {
		panic(err)
	}

	existing := MatchWordInDictionary(strings.ToUpper(word))

	if existing == true {
		json.NewEncoder(w).Encode(Score(word))
	}
}

func MatchWordInDictionary(word string) (existing bool) {
	words, err := ReadWords("dictionaries/words_pl.txt")
	if err != nil {
		return false
	}
	existing, index := inArray(word, words)
	if index == 0 {
		return false
	}
	return existing
}
