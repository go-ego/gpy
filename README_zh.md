
# gpy

[![Build Status](https://github.com/go-ego/gpy/workflows/Go/badge.svg)](https://github.com/go-ego/gpy/commits/master)
[![golangci-lint](https://img.shields.io/badge/linted%20by-golangci--lint-brightgreen)](https://golangci-lint.run/)
[![codecov](https://codecov.io/gh/go-ego/gpy/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ego/gpy)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-ego/gpy.svg)](https://pkg.go.dev/github.com/go-ego/gpy)

汉语拼音转换工具 Go 版。

[English](https://github.com/go-ego/gpy/blob/master/README.md) | 简体中文

## 安装

需要 Go 1.24 或更高版本。

```
go get github.com/go-ego/gpy
```

### 安装命令行工具

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

参数：

| 参数 | 说明 |
| --- | --- |
| `-s` | 拼音风格：`zhao`、`zh4ao`（默认）、`zha4o`、`zhao4`、`zh`、`z`、`ao`、`4ao`、`a4o`、`ao4` |
| `-e` | 启用多音字模式 |
| `-p` | 启用短语（分词）模式 |
| `-d` | 间隔符号 |
| `-v` | 启用详细输出模式 |

`-s` 和 `-d` 也可以通过 `PINYIN_S` 和 `PINYIN_D` 环境变量指定。

## 使用

- 简单汉字转换成拼音，可以使用 `github.com/go-ego/gpy`。
- 中文词句转换成拼音，结果更准确，可以使用 `github.com/go-ego/gpy/phrase`。

### 拼音风格

| 风格 | 示例（中国） |
| --- | --- |
| `gpy.Normal`（默认） | `zhong guo` |
| `gpy.Tone` | `zhōng guó` |
| `gpy.Tone2` | `zho1ng guo2` |
| `gpy.Tone3` | `zhong1 guo2` |
| `gpy.Initials` | `zh g` |
| `gpy.FirstLetter` | `z g` |
| `gpy.Finals` | `ong uo` |
| `gpy.FinalsTone` | `ōng uó` |
| `gpy.FinalsTone2` | `o1ng uo2` |
| `gpy.FinalsTone3` | `ong1 uo2` |

### gpy 使用示例

使用 `github.com/go-ego/gpy` 将中文单词逐字转换为拼音（不分词）。

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

	// 默认
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong] [guo] [hua]]

	// 包含声调
	a.Style = gpy.Tone
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhōng] [guó] [huà]]

	// 声调用数字表示（在韵母之后）
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng] [guo2] [hua4]]

	// 声调用数字表示（在拼音之后）
	a.Style = gpy.Tone3
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong1] [guo2] [hua4]]

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

	fmt.Println(gpy.Py(hans))
	// zhong guo hua
}
```

### gpy/phrase 使用示例

使用 `github.com/go-ego/gpy/phrase` 将中文句子转换为拼音。

- 基于 [gse](https://github.com/go-ego/gse) 分词，句子拼音转换更准确
- 支持自定义分词词典
- 支持自定义词语拼音
- 输出风格由 `phrase.Option` 设置（默认 `gpy.Normal` 风格）

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
	// 使用内嵌的 gse 分词词典
	phraseExampleWithEmbedDict()
	// 使用从词典文件加载的自定义分词器
	phraseExampleWithFileDict()
	// 使用 gse 默认词典文件
	phraseExampleWithDefaultDict()
}

func phraseExampleWithEmbedDict() {
	// 加载内嵌的 gse 词典
	if err := phrase.LoadGseDictEmbed("zh"); err != nil {
		log.Fatal(err)
	}
	// 将中文段落转换为拼音
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
	// xi ya tu dou hui qu; chang ye man man, winter is coming!

	// 自定义词语拼音
	// （等同于：phrase.DictAdd["都会区"] = "dū huì qū"）
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
	// xi ya tu du hui qu; chang ye man man, winter is coming!
}

// 使用自定义的 gse.Segmenter 分词词典
func phraseExampleWithFileDict() {
	// 加载 gse "zh" 词典和本地 dict.txt
	seg, err := gse.New("zh, dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase:", phrase.Paragraph(test, seg))
}

// 将 gse 默认词典文件加载到 phrase 分词器
func phraseExampleWithDefaultDict() {
	if err := phrase.LoadGseDict(); err != nil {
		log.Fatal(err)
	}
	phrase.AddDict("都会区", "dū huì qū")
	fmt.Println("gpy phrase:", phrase.Paragraph(test))
}
```

可运行的代码见 [examples](https://github.com/go-ego/gpy/tree/master/examples)（`dict.txt` 位于 [examples/phrase](https://github.com/go-ego/gpy/tree/master/examples/phrase)）。

## 相关项目

* [hotoo/pinyin](https://github.com/hotoo/pinyin): 汉语拼音转换工具 Node.js/JavaScript 版。
* [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): 汉语拼音转换工具 Python 版。
* [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): 汉语拼音转换工具 Rust 版。

## 许可证

基于 MIT 许可证，源自 [go-pinyin](https://github.com/mozillazg/go-pinyin)。
