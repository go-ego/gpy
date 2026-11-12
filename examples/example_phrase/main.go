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
