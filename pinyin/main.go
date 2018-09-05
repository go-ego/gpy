package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-ego/gpy"
	"github.com/mattn/go-isatty"
)

var (
	heteronym = flag.Bool("e", false, "启用多音字模式")

	str = `指定拼音风格。可选值：Normal, Tone, Tone2, Tone3, Initials, 
	FirstLetter, Finals, FinalsTone, FinalsTone2, FinalsTone3`
	style = flag.String("s", "Tone", str)
)

func selectArgs(args gpy.Args) {
	if *heteronym {
		args.Heteronym = true
	}
	switch *style {
	case "Normal":
		args.Style = gpy.Normal
	case "Tone2":
		args.Style = gpy.Tone2
	case "Tone3":
		args.Style = gpy.Tone3
	case "Initials":
		args.Style = gpy.Initials
	case "FirstLetter":
		args.Style = gpy.FirstLetter
	case "Finals":
		args.Style = gpy.Finals
	case "FinalsTone":
		args.Style = gpy.FinalsTone
	case "FinalsTone2":
		args.Style = gpy.FinalsTone2
	case "FinalsTone3":
		args.Style = gpy.FinalsTone3
	default:
		args.Style = gpy.Tone
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
		fmt.Println("请至少输入一个汉字: pinyin [-e] [-s STYLE] HANS [HANS ...]")
		os.Exit(1)
	}

	args := gpy.NewArgs()
	selectArgs(args)

	pys := gpy.Pinyin(strings.Join(hans, ""), args)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}
