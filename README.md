# gpy

[![Build Status](https://github.com/go-ego/gpy/workflows/Go/badge.svg)](https://github.com/go-ego/gpy/commits/master)
[![CircleCI Status](https://circleci.com/gh/go-ego/gpy.svg?style=shield)](https://circleci.com/gh/go-ego/gpy)
[![Build Status](https://travis-ci.org/go-ego/gpy.svg?branch=master)](https://travis-ci.org/go-ego/gpy)<!-- [![Coverage Status](https://coveralls.io/repos/github.com/go-ego/gpy/badge.svg?branch=master)](https://coveralls.io/r/github.com/go-ego/gpy?branch=master) -->
[![codecov](https://codecov.io/gh/go-ego/gpy/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ego/gpy)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ego/gpy)](https://goreportcard.com/report/github.com/go-ego/gpy)
[![GoDoc](https://godoc.org/github.com/go-ego/gpy?status.svg)](https://godoc.org/github.com/go-ego/gpy)

汉语拼音转换工具 Go 版。

[简体中文](https://github.com/go-ego/gpy/blob/master/README_zh.md)

## Installation

```
go get -u github.com/go-ego/gpy
```

### Install CLI tool:

```
go get -u github.com/go-ego/gpy/tools/pinyin
$ pinyin 中国话
zhōng guó huà 

$ pinyin -s zhao 中国话
zhong guo hua
```

## Usage
- If you only want to convert some simple Chinese words to pinyin, you can use the `github.com/go-ego/gpy` package.
- If you want to convert Chinese sentences to pinyin with more accurate results, you can use the `github.com/go-ego/gpy/phrase` package.

### gpy example
Use the `github.com/go-ego/gpy` package to convert Chinese words to pinyin.
```go
package main

import (
	"fmt"
	"github.com/go-ego/gpy"
)

var test = `西雅图都会区; 长夜漫漫, winter is coming!`

// if you just want to get the pinyin of a word without segmenting it, you can use the following code
func main() {
	args := gpy.Args{
		Style:     gpy.Tone,
		Heteronym: true}

	py := gpy.Pinyin(test, args)
	fmt.Println("gpy:", py)

	s := gpy.ToString(py)
	fmt.Println("gpy string:", s)
}
```

```go
package main

import (
	"fmt"

	"github.com/go-ego/gpy"
)

func main() {
	hans := "中国话"

	// 默认
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong] [guo] [hua]]

	// 包含声调
	a.Style = gpy.Tone
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhōng] [guó] [huà]]

	// 声调用数字表示
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng] [guo2] [hua4]]

	// 开启多音字模式
	a = gpy.NewArgs()
	a.Heteronym = true
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong zhong] [guo] [hua]]
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [hua4]]

	fmt.Println(gpy.LazyPinyin(hans, gpy.NewArgs()))
	// [zhong guo hua]

	fmt.Println(gpy.Convert(hans, nil))
	// [[zhong] [guo] [hua]]

	fmt.Println(gpy.LazyConvert(hans, nil))
	// [zhong guo hua]
}
```

### gpy/phrase example
Use the `github.com/go-ego/gpy/phrase` package to convert Chinese sentences to pinyin.
- Based on segment - More accurate sentence pinyin conversion
- Support for custom segmentation dictionaries
- Support for custom pinyin of words
```go
package main

import (
	"fmt"
	"github.com/go-ego/gpy/phrase"
	"github.com/go-ego/gse"
)

var test = `西雅图都会区; 长夜漫漫, winter is coming!`

func main() {
	// use default embed segmentation dict
	phraseExampleWithEmbedDict()
	// use custom file segmentation dict
	phraseExampleWithFileDict1()
	// use custom file segmentation dict
	phraseExampleWithFileDict2()
}

func phraseExampleWithEmbedDict() {
	// load default gse dict
	_ = phrase.LoadGseDictEmbed("zh")
	// convert a Chinese string paragraph to pinyin
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
	// if you want to customize the pinyin of a word, you can use the following code
	//phrase.DictAdd["都会区"] = "dū huì qū"
	phrase.AddDict("都会区", "dū huì qū")
	// convert a Chinese string paragraph to pinyin with user's dict
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
}

// if you want to customize the segmentation dict, you can use the following code
func phraseExampleWithFileDict1() {
	fmt.Println("gpy phrase 1:", phrase.Paragraph(test))
	// load gse dict from file
	seg, _ := gse.New("zh, dict.txt")
	// if you want to customize the pinyin of a word, you can use the following code
	//phrase.DictAdd["都会区"] = "dū huì qū"
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase 2:", phrase.Paragraph(test, seg))
}

// if you want to customize the segmentation dict, you can also use the following code
func phraseExampleWithFileDict2() {
	fmt.Println("gpy phrase 1:", phrase.Paragraph(test))
	phrase.LoadGseDict()
	// if you want to customize the pinyin of a word, you can use the following code
	//phrase.DictAdd["都会区"] = "dū huì qū"
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase 2:", phrase.Paragraph(test))
}
```


## Related Projects

* [hotoo/pinyin](https://github.com/hotoo/pinyin): 汉语拼音转换工具 Node.js/JavaScript 版。
* [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): 汉语拼音转换工具 Python 版。
* [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): 汉语拼音转换工具 Rust 版。

## License

Under the MIT License, base on [go-pinyin](https://github.com/mozillazg/go-pinyin).
