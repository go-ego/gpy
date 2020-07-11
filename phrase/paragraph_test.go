package phrase

import (
	"testing"

	"github.com/go-ego/gse"
	"github.com/vcaesar/tt"
)

func BenchmarkParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.046ms/op
		Paragraph("西雅图太空针, The Space Nedle")
	}
}

func TestParagraph(t *testing.T) {
	expects := map[string]string{
		"西雅图太空针, The Space Nedle": "xi ya tu tai kong zhen, The Space Nedle",
		"旧金山湾金门大桥":                "jiu jin shan wan jin men da qiao",
		"纽约帝国大厦, 纽约时代广场":          "niu yue di guo da sha, niu yue shi dai guang chang",
		"伦敦泰晤士河, 大笨钟":             "lun dun tai wu shi he, da ben zhong",
		"东京都, 东京晴空塔":              "dong jing du, dong jing qing kong ta",
		"洛杉矶好莱坞":                  "luo shan ji hao lai wu",
		"巴黎埃菲尔铁塔":                 "ba li ai fei er tie ta",
		"上海外滩, 陆家嘴上海中心":           "shang hai wai tan, lu jia zui shang hai zhong xin",
		"北京八达岭长城":                 "bei jing ba da ling chang cheng",
	}

	seg := gse.New("zh, ../examples/dict.txt")
	for source, expect := range expects {
		actual := Paragraph(source, seg)
		if expect != actual {
			tt.Equal(t, expect, actual)
			break
		}
	}

}

func TestPinyin(t *testing.T) {
	seg := gse.New("zh, ../examples/dict.txt")
	WithGse(seg)

	text := "西雅图都会区, 西雅图太空针"

	AddDict("都会区", "dū huì qū")
	p := Pinyin(text)
	tt.Equal(t, "[xi ya tu du hui qu, xi ya tu tai kong zhen]", p)

	i := Initial("都会区")
	tt.Equal(t, "dhq", i)

	Cut = false
	s := seg.Trim(seg.CutAll(text))
	i += ", "
	for _, v := range s {
		i1 := Initial(v)
		i += i1 + " "
	}
	tt.Equal(t, "dhq, xyt dhq xyt tk z ", i)
	Cut = true
}
