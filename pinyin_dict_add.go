package gpy

// PinyinDictAdd pinyin dict addition
var PinyinDictAdd = map[int]string{
	0: "",
}

// AddDict add a token into pinyin dictionary.
func AddDict(text int, py string) {
	PinyinDictAdd[text] = py
}

// Remove remove a token from pinyin dictionary.
func Remove(text int) {
	delete(PinyinDictAdd, text)
}
