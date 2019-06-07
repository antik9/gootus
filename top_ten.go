package topten

import (
	"sort"
	"strings"
)

type wordOccurence struct {
	word  string
	count int
}

func TopTen(text string) []string {
	counter := make(map[string]int)

	for _, word := range strings.Fields(text) {
		if strings.Trim(word, " \n\t\r") != "" {
			counter[word] = counter[word] + 1
		}
	}

	occurences := make([]wordOccurence, len(counter))
	topTen := make([]string, 0)
	idx, limit := 0, 10

	for key, value := range counter {
		occurences[idx] = wordOccurence{key, value}
		idx++
	}

	if idx < 10 {
		limit = idx
	}

	sort.Slice(occurences, func(i, j int) bool {
		return occurences[i].count > occurences[j].count
	})

	for _, occurence := range occurences[:limit] {
		topTen = append(topTen, occurence.word)
	}

	return topTen
}
