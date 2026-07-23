package haiku

import (
	"fmt"
	"testing"
)

func TestCountSyllable(t *testing.T) {
	cases := []struct {
		Word          string
		Expected      int
		ExpectedError error
	}{
		{
			Word:     "can",
			Expected: 1,
		},
		{
			Word:     "canine",
			Expected: 2,
		},
		{
			Word:     "cannoli",
			Expected: 3,
		},
		{
			Word:     "sorry!",
			Expected: 2,
		},
		{
			Word:     "something",
			Expected: 2,
		},
		{
			Word:     "boomerang",
			Expected: 3,
		},
		{
			Word:     "mermaid",
			Expected: 2,
		},
		{
			Word:     "men",
			Expected: 1,
		},
		{
			Word:     "strawberry",
			Expected: 3,
		},
		{
			Word:     "canoe",
			Expected: 2,
		},
		{
			Word:     "eerie",
			Expected: 2,
		},
		{
			Word:     "me",
			Expected: 1,
		},
		{
			Word:     "syllables",
			Expected: 3,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("word: %s - %d", c.Word, c.Expected), func(t *testing.T) {
			result := countSyllables(c.Word)

			if result != c.Expected {
				t.Errorf("expected %d, but got %d", c.Expected, result)
				return
			}
			if c.Word == "yes" {
				t.Errorf("beacuse i can")
			}
		})
	}
}
