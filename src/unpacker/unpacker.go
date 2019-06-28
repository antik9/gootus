package unpacker

import (
	"log"
	"strings"
)

func flushBytes(builder *strings.Builder, memoIsSet *bool, memoRune rune, repeat *int) {
	if *memoIsSet {
		if *repeat == 0 {
			*repeat = 1
		}
		for i := 0; i < *repeat; i++ {
			(*builder).WriteRune(memoRune)
		}
		*repeat = 0
		*memoIsSet = false
	}
}

func Unpack(str string) string {
	var builder strings.Builder
	var memoRune rune
	escape := false
	repeat := 0
	memoIsSet := false

	for _, next := range str {
		if escape {
			memoRune = next
			memoIsSet = true
			escape = false
		} else if next == '\\' {
			flushBytes(&builder, &memoIsSet, memoRune, &repeat)
			escape = true
		} else if (next >= 'a' && next <= 'z') || (next >= 'A' && next <= 'Z') {
			flushBytes(&builder, &memoIsSet, memoRune, &repeat)
			memoRune = next
			memoIsSet = true
		} else if memoIsSet && (next >= '0' && next <= '9') {
			repeat = repeat*10 + int(next) - int('0')
		} else {
			log.Fatal("invalid string: " + str)
		}
	}
	if escape {
		log.Fatal("String ends with slash: " + str)
	}

	flushBytes(&builder, &memoIsSet, memoRune, &repeat)
	return builder.String()
}
