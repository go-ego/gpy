/*
package gpy : 汉语拼音转换工具.

Usage

	package main

	import (
		"fmt"
		"github.com/go-ego/gpy"
	)

	func main() {
		hans := "中国人"
		// 默认
		a := gpy.NewArgs()
		fmt.Println(gpy.Pinyin(hans, a))
		// [[zhong] [guo] [ren]]

		// 包含声调
		a.Style = gpy.Tone
		fmt.Println(gpy.Pinyin(hans, a))
		// [[zhōng] [guó] [rén]]

		// 声调用数字表示
		a.Style = gpy.Tone2
		fmt.Println(gpy.Pinyin(hans, a))
		// [[zho1ng] [guo2] [re2n]]

		// 开启多音字模式
		a = gpy.NewArgs()
		a.Heteronym = true
		fmt.Println(gpy.Pinyin(hans, a))
		// [[zhong zhong] [guo] [ren]]
		a.Style = gpy.Tone2
		fmt.Println(gpy.Pinyin(hans, a))
		// [[zho1ng zho4ng] [guo2] [re2n]]
	}
*/
package gpy
