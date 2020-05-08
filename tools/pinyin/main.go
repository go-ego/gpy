package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
	"github.com/mattn/go-isatty"
)

var (
	heteronym = flag.Bool("e", false, "启用多音字模式")

	str   = `指定拼音风格。可选值：zhao, zh4ao, zha4o, zhao4, zh, z, ao, 4ao, a4o, ao4`
	style = flag.String("s", "zh4ao", str)
	phr   = flag.Bool("p", false, "Use phrase")
)

func selectArgs(args gpy.Args) {
	if *heteronym {
		args.Heteronym = true
	}

	styleValues := map[string]int{
		"zhao":  gpy.Normal,
		"zh4ao": gpy.Tone,
		"zha4o": gpy.Tone2,
		"zhao4": gpy.Tone3,
		"zh":    gpy.Initials,
		"z":     gpy.FirstLetter,
		"ao":    gpy.Finals,
		"4ao":   gpy.FinalsTone,
		"a4o":   gpy.FinalsTone2,
		"ao4":   gpy.FinalsTone3,
	}

	if value, ok := styleValues[*style]; !ok {
		fmt.Fprintf(os.Stderr, "无效的拼音风格：%s\n", *style)
		os.Exit(1)
	} else {
		args.Style = value
	}
}

func main() {
	flag.Parse()
	hans := flag.Args()

	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = ioutil.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		hans = append(hans, string(stdin))
	}

	if len(hans) == 0 {
		fmt.Fprintln(os.Stderr, "请至少输入一个汉字: pinyin [-e] [-s STYLE] HANS [HANS ...]")
		os.Exit(1)
	}

	args := gpy.NewArgs()
	selectArgs(args)

	ps := strings.Join(hans, "")
	if *phr {
		phrase.Option = args

		pys := phrase.Paragraph(ps)
		fmt.Println(pys)
		return
	}

	pys := gpy.Pinyin(ps, args)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}
