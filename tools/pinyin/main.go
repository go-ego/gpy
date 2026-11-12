package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
	"github.com/go-ego/gse"
	"github.com/mattn/go-isatty"
)

func selectArgs(args gpy.Args) gpy.Args {
	if Opts.Heteronym {
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

	if value, ok := styleValues[Opts.Style]; !ok {
		fmt.Fprintf(os.Stderr, "无效的拼音风格：%s\n", Opts.Style)
		os.Exit(1)
	} else {
		args.Style = value
	}

	return args
}

func main() {
	flag.Parse()
	if Opts.Help {
		Usage()
	}
	hans := flag.Args()

	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		var err error
		if stdin, err = io.ReadAll(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if len(stdin) > 0 {
		hans = append(hans, string(stdin))
	}

	if len(hans) == 0 {
		fmt.Fprintln(os.Stderr, "请至少输入一个汉字: pinyin [-e] [-p] [-s STYLE] HANS [HANS ...]")
		os.Exit(1)
	}

	args := gpy.NewArgs()
	args = selectArgs(args)

	ps := strings.Join(hans, "")
	if Opts.Phrase {
		phrase.Option = args

		var seg gse.Segmenter
		if !Opts.Verbose {
			seg.SkipLog = true
		}

		if err := seg.LoadDict(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		pys := phrase.Paragraph(ps, seg)
		fmt.Println(pys)
		return
	}

	pys := gpy.Pinyin(ps, args)
	for _, s := range pys {
		if Opts.Capitalize {
			s = capitalize(s)
		}
		fmt.Print(strings.Join(s, ","), Opts.Delimiter)
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}

// capitalize upper-cases the first letter of each pinyin
func capitalize(pys []string) []string {
	out := make([]string, len(pys))
	for i, p := range pys {
		if p == "" {
			continue
		}
		r, n := utf8.DecodeRuneInString(p)
		out[i] = string(unicode.ToUpper(r)) + p[n:]
	}
	return out
}
