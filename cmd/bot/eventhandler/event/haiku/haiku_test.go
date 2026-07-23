package haiku

import (
	"testing"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		Name          string
		Input         string
		Expected      string
		ExpectedError error
	}{
		{
			Name:          "struck in rear",
			Input:         "I am so sorry. Something struck me in the rear. I just wound up here.",
			Expected:      "I am so sorry.\nSomething struck me in the rear.\nI just wound up here.",
			ExpectedError: nil,
		},
		{
			Name:          "remarkable oaf",
			Input:         "Five, seven, then five Syllables mark a haiku. Remarkable oaf.",
			Expected:      "Five, seven, then five\nSyllables mark a haiku.\nRemarkable oaf.",
			ExpectedError: nil,
		},
		{
			Name:          "not an oaf",
			Input:         "They call me Sokka, that is in the Water Tribe. I am not an oaf.",
			Expected:      "They call me Sokka,\nthat is in the Water Tribe.\nI am not an oaf.",
			ExpectedError: nil,
		},
		{
			Name:          "tall",
			Input:         "Tittering monkey, In the spring he climbs treetops And thinks himself tall.",
			Expected:      "Tittering monkey,\nIn the spring he climbs treetops\nAnd thinks himself tall.",
			ExpectedError: nil,
		},
		{
			Name:          "not so hard",
			Input:         "You think you're so smart, with your fancy little words, this is not so hard.",
			Expected:      "You think you're so smart,\nwith your fancy little words,\nthis is not so hard.",
			ExpectedError: nil,
		},
		{
			Name:          "none calls it easy",
			Input:         "Whole seasons are spent Mastering the form, the style. None calls it easy.",
			Expected:      "Whole seasons are spent\nMastering the form, the style.\nNone calls it easy.",
			ExpectedError: nil,
		},
		{
			Name:          "paddle my canoe",
			Input:         "I calls it easy. Like I paddle my canoe, I'll paddle yours too!",
			Expected:      "I calls it easy.\nLike I paddle my canoe,\nI'll paddle yours too!",
			ExpectedError: nil,
		},
		{
			Name:          "nuts and fruits",
			Input:         "There's nuts and there's fruits. In Fall the clinging plum drops Always to be squashed.",
			Expected:      "There's nuts and there's fruits.\nIn Fall the clinging plum drops\nAlways to be squashed.",
			ExpectedError: nil,
		},
		{
			Name:          "boomerang",
			Input:         "Squish, squash, sling that slang. I'm always right back at ya, like my... boomerang",
			Expected:      "Squish, squash, sling that slang.\nI'm always right back at ya,\nlike my... boomerang",
			ExpectedError: nil,
		},
		{
			Name:          "not a haiku",
			Input:         "That's right, I'm Sokka, it's pronounced with an \"okka\", young ladies, I rocked ya!",
			Expected:      "",
			ExpectedError: ErrNotAHaiku,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result, err := Format(c.Input)

			if err == nil && result != c.Expected {
				t.Errorf("expected '%s' result, but got '%s'", c.Expected, result)
				return
			}

			if c.ExpectedError != nil {
				if err != c.ExpectedError {
					t.Errorf("expected '%v' error, but got '%v", c.ExpectedError, err)
				}
				if result != c.Expected {
					t.Errorf("expected '%v' result on error, but got '%v'", c.Expected, result)
				}
			}
		})
	}
}
