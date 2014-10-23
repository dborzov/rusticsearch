package main

import (
	"strings"
	"unicode/utf8"
)

func RemoveDashesEdits(word string) []string {
	var variations []string

	cursor := word
	wsCursor := word
	for {
		undashed := strings.Replace(cursor, "-", "", 1)
		if undashed == cursor {
			return variations
		}
		cursor = undashed
		variations = append(variations, cursor)
		wsCursor = strings.Replace(wsCursor, "-", " ", 1)
		variations = append(variations, wsCursor)
	}
}

func ExtraLetterEdits(word string) (variations []string) {
	count := utf8.RuneCountInString(word)
	// ignore first and last letter, usually you dont miss them in a typo
	if count < 3 {
		return
	}
	for i := 1; i < count-1; i++ {
		variations = append(variations, word[:i]+word[i+1:])
	}
	return variations
}

func SplitWordsEdit(word string) []string {
	variations := strings.Split(word, " ")
	if len(variations) < 2 {
		return []string{}
	}
	return variations
}
