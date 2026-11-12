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
