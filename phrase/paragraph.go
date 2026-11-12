// Copyright (c) 2017 go-ego
//
// All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package phrase

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gse"
)

var (
	spacesReg = regexp.MustCompile(`[\s]+`)

	// Option set pinyin style args option
	Option = gpy.Args{
		Style:     gpy.Normal,
		Heteronym: true,
	}

	hanSymbols = map[rune]rune{
		'？': '?',
		'！': '!',
		'：': ':',
		'。': '.',
		'，': ',',
		'；': ';',
		'（': '(',
		'）': ')',
		'【': '[',
		'】': ']',
		'、': ',',
		'“': '"',
		'”': '"',
	}
)

// rune classes used to decide where spaces go
const (
	clsNone   = iota
	clsSyl    // han syllable
	clsWord   // letters, digits and other non-punctuation
	clsOpen   // opening punctuation: ( [ 《 “
	clsClose  // closing and sentence punctuation: ) ] 》 ” , . ! ? ; :
	clsSymbol // other punctuation and symbols: - # $ @ ~ ...
)

func class(r rune) int {
	switch {
	case unicode.IsSpace(r):
		return clsNone
	case unicode.Is(unicode.Han, r):
		return clsSyl
	case unicode.In(r, unicode.Ps, unicode.Pi):
		return clsOpen
	case unicode.In(r, unicode.Pe, unicode.Pf),
		strings.ContainsRune(",.!?;:，。！？；：、", r):
		return clsClose
	case unicode.IsPunct(r) || unicode.IsSymbol(r):
		return clsSymbol
	}
	return clsWord
}

// needSpace reports whether a space goes between items of class prev and next;
// spaced reports whether the input already had whitespace between them.
func needSpace(prev, next int, spaced bool) bool {
	switch {
	case prev == clsNone, prev == clsOpen, next == clsClose:
		return false
	case spaced:
		return true
	case next == clsSyl:
		return prev == clsSyl || prev == clsWord || prev == clsClose
	case prev == clsSyl:
		return next == clsWord
	}
	return false
}

// Pinyin return paragraph []string
func Pinyin(p string, segs ...gse.Segmenter) []string {
	return strings.Split(Paragraph(p, segs...), " ")
}

// Initial return pinyin initial
func Initial(p string, segs ...gse.Segmenter) (s string) {
	a := Pinyin(p, segs...)
	return Join(a)
}

// Join []string to string
func Join(a []string) (s string) {
	for _, v := range a {
		if len(v) > 0 {
			s += string([]rune(v)[0])
		}
	}

	return
}

// Paragraph convert a Chinese string paragraph to pinyin,
// including letters, numbers, symbols
func Paragraph(p string, segs ...gse.Segmenter) string {
	p = pinyinPhrase(p, segs...)

	var b strings.Builder
	prev, spaced := clsNone, false
	for _, r := range p {
		if unicode.IsSpace(r) {
			spaced = true
			continue
		}

		cls, out := class(r), string(r)
		if cls == clsSyl {
			result := gpy.HanPinyin(out, Option)
			if len(result) == 0 || len(result[0]) == 0 {
				continue
			}
			out = result[0][0]
		} else if symbol, ok := hanSymbols[r]; ok {
			out = string(symbol)
		}

		if needSpace(prev, cls, spaced) {
			b.WriteByte(' ')
		}
		b.WriteString(out)
		prev, spaced = cls, false
	}

	return b.String()
}
