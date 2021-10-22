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
	allowReg  = regexp.MustCompile(`[a-zA-Z0-9\.,\?\!;\(\)\[\]\&\=\-_@\s]`)

	// Option set pinyin style args option
	Option = gpy.Args{
		Style:     gpy.Normal,
		Heteronym: true,
	}

	hanSymbols = map[string]string{
		"？": "?",
		"！": "!",
		"：": ":",
		"。": ".",
		"，": ",",
		"；": ";",
		"（": "(",
		"）": ")",
		"【": "[",
		"】": "]",
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

	for _, r := range p {
		if unicode.Is(unicode.Han, r) {
			// Han chars
			result := gpy.HanPinyin(string(r), Option)
			if len(result) == 0 {
				continue
			}

			if len(result[0]) == 0 {
				continue
			}

			s += " " + string(result[0][0]) + " "
		} else {
			// Not han chars
			char := string(r)

			if allowReg.MatchString(char) {
				s += char
			} else {
				if hanSymbols[char] != "" {
					s += hanSymbols[char]
				}
			}
		}
	}

	// trim the two continuous spaces
	s = spacesReg.ReplaceAllString(s, " ")
	m := map[string]string{
		" ,": ",",
		" .": ".",
		" ?": "?",
		" !": "!",
		" ;": ";",
		" :": ":",
		" (": "(",
		") ": ")",
		"[ ": "[",
		" ]": "]",
	}

	for k, v := range m {
		s = strings.Replace(s, k, v, -1)
	}

	s = strings.TrimSpace(s)
	return
}
