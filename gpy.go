package gpy

import (
	"strings"
	"unicode"

	"github.com/go-ego/gse"
)

// Meta
const (
	// Version get the gpy version
	Version = "v0.30.0.105"
	// License   = "MIT"
)

// GetVersion get the version
func GetVersion() string {
	return Version
}

func args(arg ...Args) Args {
	a := NewArgs()
	if len(arg) > 0 {
		a = arg[0]
	}

	return a
}

// IsChineseChar to determine whether the Chinese string
// 判断是否为中文字符串
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// Pinyin 汉字转拼音，支持多音字模式、拼音与英文等字母混合.
func Pinyin(s string, arg ...Args) [][]string {
	a := args(arg...)

	pys := [][]string{}
	sw := gse.SplitTextToWords([]byte(s))
	for i := 0; i < len(sw); i++ {
		s1 := string([]byte(sw[i]))
		r := []rune(s1)
		if len(r) <= 1 && unicode.Is(unicode.Han, r[0]) {
			py := SinglePinyin(r[0], a)
			if len(py) > 0 {
				pys = append(pys, py)
			}
		} else {
			var pya []string
			pya = append(pya, s1)
			pys = append(pys, pya)
		}
	}

	return pys
}

// ToString trans pinyin [][]string to string
func ToString(p [][]string) (s string) {
	i := 0
	for _, p1 := range p {
		r := []rune(p1[0])[0]
		i++

		if unicode.IsLetter(r) && i > 1 {
			s += " " + p1[0]
		} else {
			if i > 1 || unicode.IsSpace(r) {
				i = 0
			}
			s += p1[0]
		}
	}

	return
}

// Py return to string pinyin
func Py(s string, a ...Args) string {
	p := Pinyin(s, a...)
	return ToString(p)
}

// SinglePinyin 把单个 `rune` 类型的汉字转换为拼音.
func SinglePinyin(r rune, a Args) []string {
	if a.Fallback == nil {
		a.Fallback = Fallback
	}

	value, ok := PinyinDict[int(r)]
	if value == "" {
		value, ok = PinyinDictAdd[int(r)]
	}

	pys := []string{}
	if ok {
		pys = strings.Split(value, ",")
	} else {
		pys = a.Fallback(r, a)
	}

	if len(pys) > 0 {
		if !a.Heteronym {
			pys = pys[:1]
		}

		return applyStyle(pys, a)
	}

	return pys
}

// HanPinyin 汉字转拼音，支持多音字模式.
func HanPinyin(s string, arg ...Args) [][]string {
	a := args(arg...)

	pys := [][]string{}
	for _, r := range s {
		py := SinglePinyin(r, a)
		if len(py) > 0 {
			pys = append(pys, py)
		}
	}
	return pys
}

// LazyPinyin 汉字转拼音，与 `Pinyin` 的区别是：
// 返回值类型不同，并且不支持多音字模式，每个汉字只取第一个音.
func LazyPinyin(s string, arg ...Args) []string {
	a := args(arg...)

	a.Heteronym = false
	pys := []string{}
	for _, v := range HanPinyin(s, a) {
		pys = append(pys, v[0])
	}
	return pys
}

// Slug join `LazyPinyin` 的返回值.
// 建议改用 https://github.com/mozillazg/go-slugify
func Slug(s string, a Args) string {
	separator := a.Separator
	return strings.Join(LazyPinyin(s, a), separator)
}

// Convert 跟 Pinyin 的唯一区别就是 a 参数可以是 nil
func Convert(s string, a *Args) [][]string {
	if a == nil {
		args := NewArgs()
		a = &args
	}
	return Pinyin(s, *a)
}

// LazyConvert 跟 LazyPinyin 的唯一区别就是 a 参数可以是 nil
func LazyConvert(s string, a *Args) []string {
	if a == nil {
		args := NewArgs()
		a = &args
	}
	return LazyPinyin(s, *a)
}
