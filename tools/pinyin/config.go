// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "pinyin" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	Style      string // Style. 指定拼音风格。可选值：zhao, zh4ao, zha4o, zhao4, zh, z, ao, 4ao, a4o, ao4
	Delimiter  string // Delimiter. 间隔符号
	Capitalize bool   // Capitalize. 启用首字符大写模式
	Heteronym  bool   // Heteronym. 启用多音字模式
	Phrase     bool   // Phrase. 启用短语
	Verbose    bool   // Verbose. 启用详细输出模式
	Help       bool   // Help. 显示帮助
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {

	// set default values for command line parameters
	flag.StringVar(&Opts.Style, "s", "zh4ao",
		"Style. 指定拼音风格。可选值：zhao, zh4ao, zha4o, zhao4, zh, z, ao, 4ao, a4o, ao4")
	flag.StringVar(&Opts.Delimiter, "d", " ",
		"Delimiter. 间隔符号")
	flag.BoolVar(&Opts.Capitalize, "c", false,
		"Capitalize. 启用首字符大写模式")
	flag.BoolVar(&Opts.Heteronym, "e", false,
		"Heteronym. 启用多音字模式")
	flag.BoolVar(&Opts.Phrase, "p", false,
		"Phrase. 启用短语")
	flag.BoolVar(&Opts.Verbose, "v", false,
		"Verbose. 启用详细输出模式")
	flag.BoolVar(&Opts.Help, "help", false,
		"Help. 显示帮助")

	exists := false
	// Now override those default values from environment variables
	if len(Opts.Style) == 0 ||
		len(os.Getenv("PINYIN_S")) != 0 {
		Opts.Style = os.Getenv("PINYIN_S")
	}
	if len(Opts.Delimiter) == 0 ||
		len(os.Getenv("PINYIN_D")) != 0 {
		Opts.Delimiter = os.Getenv("PINYIN_D")
	}
	if _, exists = os.LookupEnv("PINYIN_C"); Opts.Capitalize || exists {
		Opts.Capitalize = true
	}
	if _, exists = os.LookupEnv("PINYIN_E"); Opts.Heteronym || exists {
		Opts.Heteronym = true
	}
	if _, exists = os.LookupEnv("PINYIN_P"); Opts.Phrase || exists {
		Opts.Phrase = true
	}
	if _, exists = os.LookupEnv("PINYIN_V"); Opts.Verbose || exists {
		Opts.Verbose = true
	}
	if _, exists = os.LookupEnv("PINYIN_HELP"); Opts.Help || exists {
		Opts.Help = true
	}

}

const USAGE_SUMMARY = "  -s\tStyle. 指定拼音风格。可选值：zhao, zh4ao, zha4o, zhao4, zh, z, ao, 4ao, a4o, ao4 (PINYIN_S)\n  -d\tDelimiter. 间隔符号 (PINYIN_D)\n  -c\tCapitalize. 启用首字符大写模式 (PINYIN_C)\n  -e\tHeteronym. 启用多音字模式 (PINYIN_E)\n  -p\tPhrase. 启用短语 (PINYIN_P)\n  -v\tVerbose. 启用详细输出模式 (PINYIN_V)\n  -help\tHelp. 显示帮助 (PINYIN_HELP)\n\nDetails:\n\n"

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags..] HANS [HANS ...]\n\nFlags:\n\n",
		progname)
	fmt.Fprintf(os.Stderr, USAGE_SUMMARY)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\n-s 或 -d 也可以由 PINYIN_S 或 PINYIN_D 环境变量指定\n")
	os.Exit(0)
}
