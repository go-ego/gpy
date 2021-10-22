//go:build go1.16
// +build go1.16

package phrase

// LoadGseDictEmbed load the embed dictionary
func LoadGseDictEmbed(dict ...string) error {
	loaded = true
	return seg.LoadDictEmbed(dict...)
}
