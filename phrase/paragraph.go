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
func Paragraph(p string, segs ...gse.Segmenter) (s string) {
	p = pinyinPhrase(p, segs...)
	var b strings.Builder
	var last rune
	for _, r := range p {
		if unicode.Is(unicode.Han, r) {
			// Han chars
			if last != 0 && !isPunctOrSymbol(last) {
				b.WriteRune(' ')
			}
			result := gpy.HanPinyin(string(r), Option)
			if len(result) == 0 {
				continue
			}
			if len(result[0]) == 0 {
				continue
			}
			b.WriteString(result[0][0])
		} else if symbol, ok := hanSymbols[r]; ok {
			// Han symbols
			b.WriteRune(symbol)
		} else {
			// Ohter
			b.WriteRune(r)
		}
		last = r
	}
	s = b.String()

	// trim the two continuous spaces
	s = spacesReg.ReplaceAllString(s, " ")

	s = strings.TrimSpace(s)
	return
}
