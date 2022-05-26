package tool

import (
	"bytes"
	"github.com/mmcdole/gofeed"
	"text/template"
)

func ParseTemplate(text string, item gofeed.Item) []byte {
	// 根据指定模版文本生成handler
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		panic(err)
	}
	// 模版渲染，并赋值给变量
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, item); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
