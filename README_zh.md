# gpy

[![Build Status](https://github.com/go-ego/gpy/workflows/Go/badge.svg)](https://github.com/go-ego/gpy/commits/master)
[![CircleCI Status](https://circleci.com/gh/go-ego/gpy.svg?style=shield)](https://circleci.com/gh/go-ego/gpy)
[![Build Status](https://travis-ci.org/go-ego/gpy.svg?branch=master)](https://travis-ci.org/go-ego/gpy)<!-- [![Coverage Status](https://coveralls.io/repos/gtihub.com/go-ego/gpy/badge.svg?branch=master)](https://coveralls.io/r/github.com/go-ego/gpy?branch=master) -->
[![codecov](https://codecov.io/gh/go-ego/gpy/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ego/gpy)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ego/gpy)](https://goreportcard.com/report/github.com/go-ego/gpy)
[![GoDoc](https://godoc.org/github.com/go-ego/gpy?status.svg)](https://godoc.org/github.com/go-ego/gpy)

汉语拼音转换工具 Go 版。


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


## Related Projects

* [hotoo/pinyin](https://github.com/hotoo/pinyin): 汉语拼音转换工具 Node.js/JavaScript 版。
* [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): 汉语拼音转换工具 Python 版。
* [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): 汉语拼音转换工具 Rust 版。


## License

Under the MIT License, base on [go-pinyin](https://github.com/mozillazg/go-pinyin).
