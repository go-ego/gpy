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

	phrase.LoadGseDict()
	go func() {
		fmt.Println("gpy phrase1:", phrase.Paragraph(test))
	}()
	fmt.Println("gpy phrase2:", phrase.Paragraph(test))

	seg := gse.New("zh, dict.txt")
	phrase.DictAdd["都会区"] = "dū huì qū"
	fmt.Println("gpy phrase:", phrase.Paragraph(test, seg))
}
