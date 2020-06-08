package phrase

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gse"
)

var (
	splacesRegexp    = regexp.MustCompile(`[\s]+`)
	allowCharsRegexp = regexp.MustCompile(`[a-zA-Z0-9\.,\?\!;\(\)\[\]\&\=\-_@\s]`)

	// Option set pinyin style args option
	Option = gpy.Args{
		Style:     gpy.Normal,
		Heteronym: true,
	}

	hansSymbols = map[string]string{
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
			// Other chars
			char := string(r)

			if allowCharsRegexp.MatchString(char) {
				s += char
			} else {
				if hansSymbols[char] != "" {
					s += hansSymbols[char]
				}
			}
		}
	}

	// 去掉连续两个空格
	s = splacesRegexp.ReplaceAllString(s, " ")
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
