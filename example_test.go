package gpy_test

import (
	"fmt"

	"github.com/go-ego/gpy"
)

var hans = "中国话"

func ExampleConvert() {
	fmt.Println("default:", gpy.Convert(hans, nil))
	// Output: default: [[zhong] [guo] [hua]]
}

func ExamplePinyin_default() {
	a := gpy.NewArgs()
	fmt.Println("default:", gpy.Pinyin(hans, a))
	// Output: default: [[zhong] [guo] [hua]]
}

func ExampleHanPinyin_default() {
	a := gpy.NewArgs()
	fmt.Println("default:", gpy.HanPinyin(hans, a))
	// Output: default: [[zhong] [guo] [hua]]
}

func ExamplePinyin_normal() {
	a := gpy.NewArgs()
	a.Style = gpy.Normal
	fmt.Println("Normal:", gpy.Pinyin(hans, a))
	// Output: Normal: [[zhong] [guo] [hua]]
}

func ExamplePinyin_tone() {
	a := gpy.NewArgs()
	a.Style = gpy.Tone
	fmt.Println("Tone:", gpy.Pinyin(hans, a))
	// Output: Tone: [[zhōng] [guó] [huà]]
}

func ExamplePinyin_tone2() {
	a := gpy.NewArgs()
	a.Style = gpy.Tone2
	fmt.Println("Tone2:", gpy.Pinyin(hans, a))
	// Output: Tone2: [[zho1ng] [guo2] [hua4]]
}

func ExamplePinyin_initials() {
	a := gpy.NewArgs()
	a.Style = gpy.Initials
	fmt.Println("Initials:", gpy.Pinyin(hans, a))
	// Output: Initials: [[zh] [g] [h]]
}

func ExamplePinyin_firstLetter() {
	a := gpy.NewArgs()
	a.Style = gpy.FirstLetter
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[z] [g] [h]]
}

func ExamplePinyin_finals() {
	a := gpy.NewArgs()
	a.Style = gpy.Finals
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[ong] [uo] [ua]]
}

func ExamplePinyin_finalsTone() {
	a := gpy.NewArgs()
	a.Style = gpy.FinalsTone
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[ōng] [uó] [uà]]
}

func ExamplePinyin_finalsTone2() {
	a := gpy.NewArgs()
	a.Style = gpy.FinalsTone2
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[o1ng] [uo2] [ua4]]
}

func ExamplePinyin_heteronym() {
	a := gpy.NewArgs()
	a.Heteronym = true
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// Output: [[zho1ng zho4ng] [guo2] [hua4]]
}

func ExamplePinyin_fallbackCustom1() {
	hans := "中国话abc"
	a := gpy.NewArgs()
	a.Fallback = func(r rune, a gpy.Args) []string {
		return []string{string(r + 1)}
	}
	fmt.Println(gpy.HanPinyin(hans, a))
	// Output: [[zhong] [guo] [hua] [b] [c] [d]]
}

func ExamplePinyin_fallbackCustom2() {
	hans := "中国话アイウ"
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
	// Output: [[zhong] [guo] [hua] [a] [i] [u]]
}

func ExampleLazyPinyin() {
	a := gpy.NewArgs()
	fmt.Println(gpy.LazyPinyin(hans, a))
	// Output: [zhong guo hua]
}

func ExampleSlug() {
	a := gpy.NewArgs()
	fmt.Println(gpy.Slug(hans, a))
	// Output: zhong-guo-hua
}
