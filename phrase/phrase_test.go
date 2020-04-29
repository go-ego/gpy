package phrase

import (
	"strings"
	"testing"

	"github.com/vcaesar/tt"
)

func Benchmark_pinyinPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.016ms/op
		pinyinPhrase("西雅图太空针, The Space Nedle")
	}
}

func Test_pinyinPhrase(t *testing.T) {
	expects := map[string]string{
		"西雅图太空针, The Space Nedle": "西雅图 tai kong 针, The Space Nedle",
		"旧金山湾金门大桥":                "旧金山湾金门 da qiao",
		"纽约帝国大厦, 纽约时代广场":          "纽约帝国 da sha , 纽约时代 guang chang",
		"伦敦泰晤士河, 大笨钟":             "lun dun 泰晤士河, 大笨钟",
		"东京都, 东京晴空塔":              "东 jing du , 东京 qing kong 塔",
		"洛杉矶好莱坞":                  "洛杉矶 hao lai wu",
		"巴黎埃菲尔铁塔":                 "巴黎 ai fei er tie ta",
		"上海外滩, 陆家嘴上海中心":           "shang hai 外滩, 陆家嘴 shang hai 中心",
		"北京八达岭长城":                 "bei jing ba da ling chang cheng",
	}

	for source, expect := range expects {
		actual := splacesRegexp.ReplaceAllString(strings.TrimSpace(
			pinyinPhrase(source)), " ")
		if expect != actual {
			tt.Equal(t, expect, actual)
			break
		}
	}
}
