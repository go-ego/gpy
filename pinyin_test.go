package gpy

import (
	"testing"

	"github.com/vcaesar/tt"
)

var hans = "中国话"

type (
	pinyinFunc func(string, ...Args) [][]string

	testCase struct {
		args   Args
		result [][]string
	}
)

func testPinyinTool(t *testing.T, s string, d []testCase, f pinyinFunc) {
	for _, tc := range d {
		v := f(s, tc.args)
		tt.DEqual(t, tc.result, v)
	}
}

func TestPinyin(t *testing.T) {
	var pyStr = [][]string{
		{"zhong"},
		{"guo"},
		{"hua"},
	}

	testData := []testCase{
		// default
		{
			Args{Style: Normal},
			pyStr,
		},
		// default
		{
			NewArgs(),
			pyStr,
		},
		// Normal
		{
			Args{Style: Normal},
			pyStr,
		},
		// Tone
		{
			Args{Style: Tone},
			[][]string{
				{"zhōng"},
				{"guó"},
				{"huà"},
			},
		},
		// Tone2
		{
			Args{Style: Tone2},
			[][]string{
				{"zho1ng"},
				{"guo2"},
				{"hua4"},
			},
		},
		// Tone3
		{
			Args{Style: Tone3},
			[][]string{
				{"zhong1"},
				{"guo2"},
				{"hua4"},
			},
		},
		// Initials
		{
			Args{Style: Initials},
			[][]string{
				{"zh"},
				{"g"},
				{"h"},
			},
		},
		// FirstLetter
		{
			Args{Style: FirstLetter},
			[][]string{
				{"z"},
				{"g"},
				{"h"},
			},
		},
		// Finals
		{
			Args{Style: Finals},
			[][]string{
				{"ong"},
				{"uo"},
				{"ua"},
			},
		},
		// FinalsTone
		{
			Args{Style: FinalsTone},
			[][]string{
				{"ōng"},
				{"uó"},
				{"uà"},
			},
		},
		// FinalsTone2
		{
			Args{Style: FinalsTone2},
			[][]string{
				{"o1ng"},
				{"uo2"},
				{"ua4"},
			},
		},
		// FinalsTone3
		{
			Args{Style: FinalsTone3},
			[][]string{
				{"ong1"},
				{"uo2"},
				{"ua4"},
			},
		},
		// Heteronym
		{
			Args{Heteronym: true},
			[][]string{
				{"zhong", "zhong"},
				{"guo"},
				{"hua"},
			},
		},
	}

	testPinyinTool(t, hans, testData, Pinyin)

	// 测试不是多音字的 Heteronym
	hans := "你"
	testData = []testCase{
		{
			Args{},
			[][]string{
				{"ni"},
			},
		},
		{
			Args{Heteronym: true},
			[][]string{
				{"ni"},
			},
		},
	}
	testPinyinTool(t, hans, testData, Pinyin)
}

func TestNoneHans(t *testing.T) {
	s := "abc"
	v := HanPinyin(s, NewArgs())
	value := [][]string{}
	tt.DEqual(t, value, v)
}

func TestNone(t *testing.T) {
	s := "abc"
	v := Pinyin(s, NewArgs())
	value := [][]string{{"abc"}}
	tt.DEqual(t, value, v)
}

func TestLazyPinyin(t *testing.T) {
	v := LazyPinyin(hans, Args{})
	value := []string{"zhong", "guo", "hua"}
	tt.DEqual(t, value, v)

	hans := "中国话abc"
	v = LazyPinyin(hans, Args{})
	value = []string{"zhong", "guo", "hua"}
	tt.DEqual(t, value, v)
}

func TestSlug(t *testing.T) {
	v := Slug(hans, Args{})
	value := "zhongguohua"
	tt.Equal(t, value, v)

	v = Slug(hans, Args{Separator: ","})
	value = "zhong,guo,hua"
	tt.Equal(t, value, v)

	a := NewArgs()
	v = Slug(hans, a)
	value = "zhong-guo-hua"
	tt.Equal(t, value, v)

	hans := "中国话abc，,中"
	v = Slug(hans, a)
	value = "zhong-guo-hua-zhong"
	tt.Equal(t, value, v)
}

func TestFinal(t *testing.T) {
	value := "an"
	v := final("an")
	tt.Equal(t, value, v)
}

func TestFallback(t *testing.T) {
	hans := "中国话abc"
	testData := []testCase{
		// default
		{
			NewArgs(),
			[][]string{
				{"zhong"},
				{"guo"},
				{"hua"},
			},
		},
		// custom
		{
			Args{
				Fallback: func(r rune, a Args) []string {
					return []string{"la"}
				},
			},
			[][]string{
				{"zhong"},
				{"guo"},
				{"hua"},
				{"la"},
				{"la"},
				{"la"},
			},
		},
		// custom
		{
			Args{
				Heteronym: true,
				Fallback: func(r rune, a Args) []string {
					return []string{"la", "wo"}
				},
			},
			[][]string{
				{"zhong", "zhong"},
				{"guo"},
				{"hua"},
				{"la", "wo"},
				{"la", "wo"},
				{"la", "wo"},
			},
		},
	}
	testPinyinTool(t, hans, testData, HanPinyin)
}

type testItem struct {
	hans   string
	args   Args
	result [][]string
}

func testPinyinToolUpdate(t *testing.T, d []testItem, f pinyinFunc, call string) {
	for _, tc := range d {
		v := f(tc.hans, tc.args)
		tt.DEqual(t, tc.result, v, "", call)
	}
}

func TestUpdated(t *testing.T) {
	testData := []testItem{
		// 误把 yu 放到声母列表了
		{"鱼", Args{Style: Tone2}, [][]string{{"yu2"}}},
		{"鱼", Args{Style: Tone3}, [][]string{{"yu2"}}},
		{"鱼", Args{Style: Finals}, [][]string{{"v"}}},
		{"雨", Args{Style: Tone2}, [][]string{{"yu3"}}},
		{"雨", Args{Style: Tone3}, [][]string{{"yu3"}}},
		{"雨", Args{Style: Finals}, [][]string{{"v"}}},
		{"元", Args{Style: Tone2}, [][]string{{"yua2n"}}},
		{"元", Args{Style: Tone3}, [][]string{{"yuan2"}}},
		{"元", Args{Style: Finals}, [][]string{{"van"}}},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		{"呀", Args{Style: Initials}, [][]string{{""}}},
		{"呀", Args{Style: Tone2}, [][]string{{"ya"}}},
		{"呀", Args{Style: Tone3}, [][]string{{"ya"}}},
		{"呀", Args{Style: Finals}, [][]string{{"ia"}}},
		{"无", Args{Style: Initials}, [][]string{{""}}},
		{"无", Args{Style: Tone2}, [][]string{{"wu2"}}},
		{"无", Args{Style: Tone3}, [][]string{{"wu2"}}},
		{"无", Args{Style: Finals}, [][]string{{"u"}}},
		{"衣", Args{Style: Tone2}, [][]string{{"yi1"}}},
		{"衣", Args{Style: Tone3}, [][]string{{"yi1"}}},
		{"衣", Args{Style: Finals}, [][]string{{"i"}}},
		{"万", Args{Style: Tone2}, [][]string{{"wa4n"}}},
		{"万", Args{Style: Tone3}, [][]string{{"wan4"}}},
		{"万", Args{Style: Finals}, [][]string{{"uan"}}},
		// ju, qu, xu 的韵母应该是 v
		{"具", Args{Style: FinalsTone}, [][]string{{"ǜ"}}},
		{"具", Args{Style: FinalsTone2}, [][]string{{"v4"}}},
		{"具", Args{Style: FinalsTone3}, [][]string{{"v4"}}},
		{"具", Args{Style: Finals}, [][]string{{"v"}}},
		{"取", Args{Style: FinalsTone}, [][]string{{"ǚ"}}},
		{"取", Args{Style: FinalsTone2}, [][]string{{"v3"}}},
		{"取", Args{Style: FinalsTone3}, [][]string{{"v3"}}},
		{"取", Args{Style: Finals}, [][]string{{"v"}}},
		{"徐", Args{Style: FinalsTone}, [][]string{{"ǘ"}}},
		{"徐", Args{Style: FinalsTone2}, [][]string{{"v2"}}},
		{"徐", Args{Style: FinalsTone3}, [][]string{{"v2"}}},
		{"徐", Args{Style: Finals}, [][]string{{"v"}}},
		// # ń
		{"嗯", Args{Style: Normal}, [][]string{{"n"}}},
		{"嗯", Args{Style: Tone}, [][]string{{"ń"}}},
		{"嗯", Args{Style: Tone2}, [][]string{{"n2"}}},
		{"嗯", Args{Style: Tone3}, [][]string{{"n2"}}},
		{"嗯", Args{Style: Initials}, [][]string{{""}}},
		{"嗯", Args{Style: FirstLetter}, [][]string{{"n"}}},
		{"嗯", Args{Style: Finals}, [][]string{{"n"}}},
		{"嗯", Args{Style: FinalsTone}, [][]string{{"ń"}}},
		{"嗯", Args{Style: FinalsTone2}, [][]string{{"n2"}}},
		{"嗯", Args{Style: FinalsTone3}, [][]string{{"n2"}}},
		// # ḿ  \u1e3f  U+1E3F
		{"呣", Args{Style: Normal}, [][]string{{"m"}}},
		{"呣", Args{Style: Tone}, [][]string{{"ḿ"}}},
		{"呣", Args{Style: Tone2}, [][]string{{"m2"}}},
		{"呣", Args{Style: Tone3}, [][]string{{"m2"}}},
		{"呣", Args{Style: Initials}, [][]string{{""}}},
		{"呣", Args{Style: FirstLetter}, [][]string{{"m"}}},
		{"呣", Args{Style: Finals}, [][]string{{"m"}}},
		{"呣", Args{Style: FinalsTone}, [][]string{{"ḿ"}}},
		{"呣", Args{Style: FinalsTone2}, [][]string{{"m2"}}},
		{"呣", Args{Style: FinalsTone3}, [][]string{{"m2"}}},
		// 去除 0
		{"啊", Args{Style: Tone2}, [][]string{{"a"}}},
		{"啊", Args{Style: Tone3}, [][]string{{"a"}}},
		{"侵略", Args{Style: Tone2}, [][]string{{"qi1n"}, {"lve4"}}},
		{"侵略", Args{Style: FinalsTone2}, [][]string{{"i1n"}, {"ve4"}}},
		{"侵略", Args{Style: FinalsTone3}, [][]string{{"in1"}, {"ve4"}}},
	}
	testPinyinToolUpdate(t, testData, Pinyin, "pinyin_test.go:345")
}

func TestConvert(t *testing.T) {
	v := Convert(hans, nil)
	value := [][]string{{"zhong"}, {"guo"}, {"hua"}}
	tt.DEqual(t, value, v)

	a := NewArgs()
	v = Convert(hans, &a)
	tt.DEqual(t, value, v)
}

func TestLazyConvert(t *testing.T) {
	v := LazyConvert(hans, nil)
	value := []string{"zhong", "guo", "hua"}
	tt.DEqual(t, value, v)

	a := NewArgs()
	v = LazyConvert(hans, &a)
	tt.DEqual(t, value, v)
}

func TestPy(t *testing.T) {
	v1 := "zhong guo hua"
	v := Py(hans)
	tt.Equal(t, v1, v)
}
