package main

import (
	"bufio"
	"os"
	"reflect"
)

var wordsDict []string

func ReadWords(path string) ([]string, error) {
	if len(wordsDict) > 0 {
		return wordsDict, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordsDict = append(wordsDict, scanner.Text())
	}

	return wordsDict, scanner.Err()
}

func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
