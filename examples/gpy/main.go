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
