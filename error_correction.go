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

func ExtraLetterEdits(word string) []string {
	var variations []string
	for i := 1; i <= utf8.RuneCountInString(word); i++ {
		variations = append(variations, word[:i-1]+word[i:])
	}
	return variations
}
