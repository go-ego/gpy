package phrase

import (
	"strings"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gse"
)

var (
	seg    gse.Segmenter
	loaded bool
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
		seg = gse.New()
		loaded = true
	}
	return seg.CutAll(s)
}

func pinyinPhrase(s string, segs ...gse.Segmenter) string {
	words := cutWords(s, segs...)
	for _, word := range words {
		match := phraseDict[word]
		if match == "" {
			match = DictAdd[word]
		}

		match = gpy.ToFixed(match, Option)
		if match != "" {
			s = strings.Replace(s, word, " "+match+" ", 1)
		}
	}

	return s
}
