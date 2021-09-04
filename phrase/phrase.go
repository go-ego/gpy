package phrase

import (
	"strings"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gse"
)

var (
	seg    gse.Segmenter
	loaded bool

	// Cut set pinyinPhrase cut
	Cut = true
)

// LoadGseDict load the user's gse dict
func LoadGseDict(files ...string) error {
	loaded = true
	return seg.LoadDict(files...)
}

// WithGse register gse segmenter
func WithGse(segs gse.Segmenter) {
	seg = segs
	loaded = true
}

// CutWord cut the string word
func CutWord(s string) []string {
	return seg.CutAll(s)
}

func cutWords(s string, segs ...gse.Segmenter) []string {
	if len(segs) > 0 {
		seg = segs[0]
		loaded = true
	}

	if !loaded {
		seg, _ = gse.New()
		loaded = true
	}
	return seg.CutAll(s)
}

// Match match word pinyin
func Match(word string) string {
	match := phraseDict[word]
	if match == "" {
		match = DictAdd[word]
	}

	match = gpy.ToFixed(match, Option)
	return match
}

func matchs(s, word string) string {
	match := Match(word)
	if match != "" {
		s = strings.Replace(s, word, " "+match+" ", 1)
	}

	return s
}

func pinyinPhrase(s string, segs ...gse.Segmenter) string {
	if Cut {
		words := cutWords(s, segs...)
		for _, word := range words {
			s = matchs(s, word)
		}

		return s
	}

	s = matchs(s, s)
	return s
}
