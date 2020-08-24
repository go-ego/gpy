## Usage

```go
package main

import (
	"fmt"

	"github.com/go-ego/gse"
	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
)

var test = `那里湖面总是澄清, 那里空气充满宁静; 西雅图都会区`

func main() {
	args := gpy.Args{
		Style:     gpy.Tone,
		Heteronym: true}

	py := gpy.Pinyin(test, args)
	fmt.Println("gpy:", py)

	s := gpy.ToString(py)
	fmt.Println("gpy string:", s)

	phrase.LoadGseDict()
	
	seg := gse.New("zh, dict.txt")
	phrase.AddDict("都会区", "dū huì qū")

	fmt.Println("gpy phrase:", phrase.Paragraph(test, seg))
	fmt.Println("pinyin: ", phrase.Pinyin(test))
	fmt.Println("Initial: ", phrase.Initial("都会区"))
}
```
