
# gpy

[![Build Status](https://github.com/go-ego/gpy/workflows/Go/badge.svg)](https://github.com/go-ego/gpy/commits/master)
[![golangci-lint](https://img.shields.io/badge/linted%20by-golangci--lint-brightgreen)](https://golangci-lint.run/)
[![codecov](https://codecov.io/gh/go-ego/gpy/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ego/gpy)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-ego/gpy.svg)](https://pkg.go.dev/github.com/go-ego/gpy)

Chinese to Hanyu Pinyin conversion in Go (汉语拼音转换工具 Go 版).

English | [简体中文](https://github.com/go-ego/gpy/blob/master/README_zh.md)

## Installation

Requires Go 1.24 or later.

```
go get github.com/go-ego/gpy
```

### Install the CLI tool

```
go install github.com/go-ego/gpy/tools/pinyin@latest
```

```
$ pinyin 中国话
zhōng guó huà

$ pinyin -s zhao 中国话
zhong guo hua

$ pinyin -p 西雅图都会区
xī yǎ tú dōu huì qū
```

Flags:

| Flag | Description |
| --- | --- |
| `-s` | Pinyin style: `zhao`, `zh4ao` (default), `zha4o`, `zhao4`, `zh`, `z`, `ao`, `4ao`, `a4o`, `ao4` |
| `-e` | Enable heteronym (multiple pronunciations) mode |
| `-p` | Enable phrase (word segmentation) mode |
| `-d` | Delimiter |
| `-v` | Verbose output |

`-s` and `-d` can also be set with the `PINYIN_S` and `PINYIN_D` environment variables.

## Usage

- To convert simple Chinese words to pinyin, use the `github.com/go-ego/gpy` package.
- To convert Chinese sentences to pinyin with more accurate results, use the `github.com/go-ego/gpy/phrase` package.

### Styles

| Style | Example (中国) |
| --- | --- |
| `gpy.Normal` (default) | `zhong guo` |
| `gpy.Tone` | `zhōng guó` |
| `gpy.Tone2` | `zho1ng guo2` |
| `gpy.Tone3` | `zhong1 guo2` |
| `gpy.Initials` | `zh g` |
| `gpy.FirstLetter` | `z g` |
| `gpy.Finals` | `ong uo` |
| `gpy.FinalsTone` | `ōng uó` |
| `gpy.FinalsTone2` | `o1ng uo2` |
| `gpy.FinalsTone3` | `ong1 uo2` |

### gpy example

Use the `github.com/go-ego/gpy` package to convert Chinese words to pinyin, character by character, without word segmentation.

```go
package main

import (
	"fmt"

	"github.com/go-ego/gpy"
)

var test = `西雅图都会区; 长夜漫漫, winter is coming!`

func main() {
	args := gpy.Args{
		Style:     gpy.Tone,
		Heteronym: true}

	py := gpy.Pinyin(test, args)
	fmt.Println("gpy:", py)
	// gpy: [[xī] [yǎ yā yá] [tú] [dōu dū] [huì kuài] [qū ōu] [;] [ ] [zhǎng cháng] ...]

	s := gpy.ToString(py)
	fmt.Println("gpy string:", s)
	// gpy string: xī yǎ tú dōu huì qū; zhǎng yè màn màn, winter is coming!
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

	// default style
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong] [guo] [hua]]

	// with tone marks
	a.Style = gpy.Tone
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhōng] [guó] [huà]]

	// tones as numbers after the vowel
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng] [guo2] [hua4]]

	// tones as numbers at the end
	a.Style = gpy.Tone3
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong1] [guo2] [hua4]]

	// heteronym mode
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

	fmt.Println(gpy.Py(hans))
	// zhong guo hua
}
```

### gpy/phrase example

Use the `github.com/go-ego/gpy/phrase` package to convert Chinese sentences to pinyin.

- Based on [gse](https://github.com/go-ego/gse) word segmentation - more accurate sentence pinyin conversion
- Supports custom segmentation dictionaries
- Supports custom pinyin for words
- The output style is set by `phrase.Option` (default: `gpy.Normal` style)

```go
package main

import (
	"fmt"
	"log"

	"github.com/go-ego/gpy/phrase"
	"github.com/go-ego/gse"
)

var test = `西雅图都会区; 长夜漫漫, winter is coming!`

func main() {
	// use the embedded gse segmentation dict
	phraseExampleWithEmbedDict()
	// use a custom segmenter loaded from a dict file
	phraseExampleWithFileDict()
	// use the default gse dict files
	phraseExampleWithDefaultDict()
}

func phraseExampleWithEmbedDict() {
	// load the embedded gse dict
	if err := phrase.LoadGseDictEmbed("zh"); err != nil {
		log.Fatal(err)
	}
	// convert a Chinese paragraph to pinyin
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
	// xi ya tu dou hui qu; chang ye man man, winter is coming!

	// customize the pinyin of a word
	// (same as: phrase.DictAdd["都会区"] = "dū huì qū")
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
	// xi ya tu du hui qu; chang ye man man, winter is coming!
}

// customize the segmentation dict with your own gse.Segmenter
func phraseExampleWithFileDict() {
	// load the gse "zh" dict and the local dict.txt
	seg, err := gse.New("zh, dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase:", phrase.Paragraph(test, seg))
}

// load the default gse dict files into the phrase segmenter
func phraseExampleWithDefaultDict() {
	if err := phrase.LoadGseDict(); err != nil {
		log.Fatal(err)
	}
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
}
```

See [examples](https://github.com/go-ego/gpy/tree/master/examples) for runnable code (`dict.txt` is in [examples/phrase](https://github.com/go-ego/gpy/tree/master/examples/phrase)).

## Related Projects

* [hotoo/pinyin](https://github.com/hotoo/pinyin): Chinese to pinyin for Node.js/JavaScript.
* [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): Chinese to pinyin for Python.
* [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): Chinese to pinyin for Rust.

## License

Under the MIT License, based on [go-pinyin](https://github.com/mozillazg/go-pinyin).
