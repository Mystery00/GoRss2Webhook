package tool

import (
	"github.com/mmcdole/gofeed"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	tpl := `标题: {{.Title}} 描述: {{.Description}}`
	item := gofeed.Item{
		Title:       "国家卫健委",
		Description: "5月21日0—24时",
	}
	bytes := ParseTemplate(tpl, item)
	if string(bytes) != "标题: 国家卫健委 描述: 5月21日0—24时" {
		t.Error(`parseTemplate failed`)
	}
}
