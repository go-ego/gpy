// Copyright (c) 2017 go-ego
//
// All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package phrase

import (
	"strings"
	"unicode"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gse"
)

var (
	seg    gse.Segmenter
	loaded bool

	// Cut set the pinyin phrase cut
	Cut = true
)

// LoadGseDict load the user's gse dict
func LoadGseDict(files ...string) error {
	loaded = true
	return seg.LoadDict(files...)
}

// WithGse register the gse segmenter
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

// Match match the word pinyin
func Match(word string) string {
	match := phraseDict[word]
	if match == "" {
		match = DictAdd[word]
	}

	match = gpy.ToFixed(match, Option)
	return match
}

func isPunctOrSymbol(r rune) bool {
	return unicode.IsOneOf([]*unicode.RangeTable{unicode.Punct, unicode.Symbol}, r)
}

func matchs(s, word string) string {
	if match := Match(word); match != "" {
		if before, after, found := strings.Cut(s, word); found {
			if before := []rune(before); len(before) > 0 && !isPunctOrSymbol(before[len(before)-1]) {
				match = " " + match
			}
			if after := []rune(after); len(after) > 0 && !isPunctOrSymbol(after[0]) {
				match += " "
			}
			s = strings.Replace(s, word, match, 1)
		}
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

	return matchs(s, s)
}
