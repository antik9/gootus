package topten

import (
	"sort"
	"strings"
	"unicode"
)

type wordOccurence struct {
	word  string
	count int
}

func parseWords(text string) []string {
	return strings.FieldsFunc(
		text, func(r rune) bool {
			return unicode.IsSpace(r) || unicode.IsPunct(r) ||
				unicode.IsControl(r) || unicode.IsSymbol(r)
		})
}

func getWordOccurences(words []string) []wordOccurence {
	counter := make(map[string]int)
	for _, word := range words {
		counter[word] = counter[word] + 1
	}

	occurences := make([]wordOccurence, 0, len(counter))

	for key, value := range counter {
		occurences = append(occurences, wordOccurence{key, value})
	}
	return occurences
}

func sortWordOccurences(occurences []wordOccurence) []wordOccurence {
	sort.Slice(occurences, func(i, j int) bool {
		return occurences[i].count > occurences[j].count
	})
	return occurences
}

func TopN(text string, limit int) []string {
	words := parseWords(text)
	occurences := sortWordOccurences(
		getWordOccurences(words))

	if limit > len(occurences) {
		limit = len(occurences)
	}
	top := make([]string, 0, limit)

	for _, occurence := range occurences[:limit] {
		top = append(top, occurence.word)
	}
	return top
}

func TopTen(text string) []string {
	return TopN(text, 10)
}
