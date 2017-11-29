package gpy_test

import (
	"fmt"

	"github.com/go-ego/gpy"
)

func ExampleConvert() {
	hans := "中国人"
	fmt.Println("default:", gpy.Convert(hans, nil))
	// Output: default: [[zhong] [guo] [ren]]
}

func ExamplePinyin_default() {
	hans := "中国人"
	a := gpy.NewArgs()
	fmt.Println("default:", gpy.Pinyin(hans, a))
	// Output: default: [[zhong] [guo] [ren]]
}

func ExampleHanPinyin_default() {
	hans := "中国人"
	a := gpy.NewArgs()
	fmt.Println("default:", gpy.HanPinyin(hans, a))
	// Output: default: [[zhong] [guo] [ren]]
}

func ExamplePinyin_normal() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.Normal
	fmt.Println("Normal:", gpy.Pinyin(hans, a))
	// Output: Normal: [[zhong] [guo] [ren]]
}

func ExamplePinyin_tone() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.Tone
	fmt.Println("Tone:", gpy.Pinyin(hans, a))
	// Output: Tone: [[zhōng] [guó] [rén]]
}

func ExamplePinyin_tone2() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.Tone2
	fmt.Println("Tone2:", gpy.Pinyin(hans, a))
	// Output: Tone2: [[zho1ng] [guo2] [re2n]]
}

func ExamplePinyin_initials() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.Initials
	fmt.Println("Initials:", gpy.Pinyin(hans, a))
	// Output: Initials: [[zh] [g] [r]]
}

func ExamplePinyin_firstLetter() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.FirstLetter
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[z] [g] [r]]
}

func ExamplePinyin_finals() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.Finals
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[ong] [uo] [en]]
}

func ExamplePinyin_finalsTone() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.FinalsTone
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[ōng] [uó] [én]]
}

func ExamplePinyin_finalsTone2() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Style = gpy.FinalsTone2
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[o1ng] [uo2] [e2n]]
}

func ExamplePinyin_heteronym() {
	hans := "中国人"
	a := gpy.NewArgs()
	a.Heteronym = true
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[zho1ng zho4ng] [guo2] [re2n]]
}

func ExamplePinyin_fallbackCustom1() {
	hans := "中国人abc"
	a := gpy.NewArgs()
	a.Fallback = func(r rune, a gpy.Args) []string {
		return []string{string(r + 1)}
	}
	fmt.Println(gpy.HanPinyin(hans, a))
	// Output: [[zhong] [guo] [ren] [b] [c] [d]]
}

func ExamplePinyin_fallbackCustom2() {
	hans := "中国人アイウ"
	a := gpy.NewArgs()
	a.Fallback = func(r rune, a gpy.Args) []string {
		data := map[rune][]string{
			'ア': {"a"},
			'イ': {"i"},
			'ウ': {"u"},
		}
		s, ok := data[r]
		if ok {
			return s
		} else {
			return []string{}
		}
	}
	fmt.Println(gpy.HanPinyin(hans, a))
	// Output: [[zhong] [guo] [ren] [a] [i] [u]]
}

func ExampleLazyPinyin() {
	hans := "中国人"
	a := gpy.NewArgs()
	fmt.Println(gpy.LazyPinyin(hans, a))
	// Output: [zhong guo ren]
}

func ExampleSlug() {
	hans := "中国人"
	a := gpy.NewArgs()
	fmt.Println(gpy.Slug(hans, a))
	// Output: zhong-guo-ren
}
