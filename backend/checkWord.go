package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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
	fmt.Println(existing)
	json.NewEncoder(w).Encode(existing)
}

func MatchWordInDictionary(word string) (existing bool) {
	words, err := ReadWords("dictionaries/words.txt")
	if err != nil {
		return false
	}
	existing, index := inArray(word, words)
	if index == 0 {
		return false
	}
	return existing
}
