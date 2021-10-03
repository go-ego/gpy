package phrase

/*
	go pinyin addition dictionary
*/

// DictAdd phrase dict addition map
var DictAdd = map[string]string{
	"宿舍": "sù shè",
	"不薄": "bù báo",
	"打折": "dǎ zhé",
	"着手": "zhuó shǒu",
	"着眼": "zhuó yǎn",
	"着重": "zhuó zhòng",
}

// AddDict add a token into phrase dictionary.
func AddDict(text, py string) {
	DictAdd[text] = py
}

// Remove remove a token from phrase dictionary.
func Remove(text string) {
	delete(DictAdd, text)
}
