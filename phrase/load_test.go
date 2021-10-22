//go:build go1.16
// +build go1.16

package phrase

import (
	"testing"

	"github.com/vcaesar/tt"
)

func TestLoadDict(t *testing.T) {
	err := LoadGseDictEmbed("zh_t")
	tt.Nil(t, err)

	text := "尼亚加拉大瀑布, 多伦多电视塔天空巨蛋"
	p := Pinyin(text)
	tt.Equal(t, "[ni ya jia la da pu bu, duo lun duo dian shi ta tian kong ju dan]", p)

	err = LoadGseDict("zh_t")
	tt.Nil(t, err)
}
