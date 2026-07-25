package haiku

import (
	"fmt"
	"strings"
)

func newWord(content string) *word {
	return &word{
		Content:       content,
		SyllableCount: countSyllables(content),
	}
}

type word struct {
	Content       string
	SyllableCount int
}

func newLine() *line {
	return &line{words: []*word{}}
}

type line struct {
	words []*word
}

func (l *line) Append(content string) {
	l.words = append(l.words, newWord(content))
}

func (l *line) Format() string {
	results := []string{}
	for _, w := range l.words {
		results = append(results, w.Content)
	}
	return strings.Join(results, " ")
}

func (l *line) SyllableCount() int {
	total := 0
	for _, w := range l.words {

		total += w.SyllableCount
	}
	return total
}

func newPhrase() *phrase {
	return &phrase{
		line1: newLine(),
		line2: newLine(),
		line3: newLine(),
	}
}

type phrase struct {
	line1 *line
	line2 *line
	line3 *line
}

func (p *phrase) AddWord(word string) {
	if p.line1.SyllableCount() < 5 {
		p.line1.Append(word)
	} else if p.line2.SyllableCount() < 7 {
		p.line2.Append(word)
	} else {
		p.line3.Append(word)
	}
}

func (p *phrase) Format() string {
	return fmt.Sprintf("%s\n%s\n%s", p.line1.Format(), p.line2.Format(), p.line3.Format())
}

func (p *phrase) IsValid() bool {
	return p.line1.SyllableCount() == 5 && p.line2.SyllableCount() == 7 && p.line3.SyllableCount() == 5
}

var ErrNotAHaiku = fmt.Errorf("not a haiku")

func Format(message string) (string, error) {
	phrase := newPhrase()

	for _, w := range strings.Fields(strings.TrimSpace(message)) {
		phrase.AddWord(w)
	}
	if !phrase.IsValid() {
		return "", ErrNotAHaiku
	}
	return phrase.Format(), nil
}
