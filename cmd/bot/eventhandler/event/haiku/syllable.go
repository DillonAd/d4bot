package haiku

import (
	"fmt"
	"slices"
	"strings"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u', 'y'}

func countSyllables(word string) int {
	vowelCount := 0
	lastCharWasVowel := false

	wrd := formatInput(word)

	for _, c := range wrd {
		if slices.Contains(vowels, c) {
			if !lastCharWasVowel {
				vowelCount++
			}
			lastCharWasVowel = true
		} else {
			lastCharWasVowel = false
		}
	}

	if hasSilentME(wrd) {
		vowelCount--
	} else if hasSilentES(wrd) || hasSilentE(wrd) {
		vowelCount--
		if vowelCount <= 0 {
			vowelCount = 1
		}
	}

	return vowelCount
}

func formatInput(input string) string {
	result := []rune{}
	for _, r := range strings.ToLower(input) {
		if 'a' <= r && r <= 'z' {
			result = append(result, r)
		}
	}
	return string(result)
}

func hasSilentME(word string) bool {
	return len(word) > 3 && strings.Contains(word, "me") && !strings.Contains(word, "mer")
}

func hasSilentES(word string) bool {
	return strings.HasSuffix(word, "es") && !hasSkippableSuffix(word, "es")
}

func hasSilentE(word string) bool {
	return strings.HasSuffix(word, "e") && !hasSkippableSuffix(word, "e")
}

var skippablePenultimateChars = append(vowels, []rune{'l'}...)

func hasSkippableSuffix(word, suffix string) bool {
	return slices.ContainsFunc(skippablePenultimateChars, func(c rune) bool {
		return strings.HasSuffix(word, fmt.Sprintf("%c%s", c, suffix))
	})
}
