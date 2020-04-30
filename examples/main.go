package main

import (
	"fmt"

	"github.com/go-ego/gse"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
)

var test = `西雅图都会区; 长夜漫漫, winter is coming!`

func main() {
	args := gpy.Args{
		Style:     gpy.Tone,
		Heteronym: true}

	py := gpy.Pinyin(test, args)
	fmt.Println("gpy:", py)

	s := gpy.ToString(py)
	fmt.Println("gpy string:", s)

	seg := gse.New("zh, dict.txt")
	fmt.Println("gpy phrase:", phrase.Paragraph(test, seg))
}
