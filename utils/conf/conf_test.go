package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

type test struct {
	TT struct {
		Enable bool `mapstructure:"enable"`
	} `mapstructure:"tt"`
	ARR []struct {
		ITEM string `mapstructure:"item"`
	} `mapstructure:"arr"`
}

func TestMarshal(t *testing.T) {
	storePath := `/tmp/GoRss2Webhook/conf_store`
	storeFileName := `conf`
	configType := `yaml`

	err := os.MkdirAll(storePath, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	tt := test{
		TT: struct {
			Enable bool `mapstructure:"enable"`
		}{
			Enable: true,
		},
		ARR: []struct {
			ITEM string `mapstructure:"item"`
		}{
			{
				ITEM: "1",
			},
			{
				ITEM: "2",
			},
		},
	}
	bs, err := yaml.Marshal(tt)
	if err != nil {
		t.Error(err)
	}
	filePath := fmt.Sprintf(`%s/%s.%s`, storePath, storeFileName, configType)
	err = ioutil.WriteFile(filePath, bs, 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal(t *testing.T) {
	storePath := `/tmp/GoRss2Webhook/conf_store`
	storeFileName := `conf`
	configType := `yaml`

	err := os.MkdirAll(storePath, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	filePath := fmt.Sprintf(`%s/%s.%s`, storePath, storeFileName, configType)
	content := `
tt:
  enable: true
arr:
- item: "1"
- item: "2"
`
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Error(err)
	}
	config := InitConfig(storePath, storeFileName, configType)
	var tt test
	err = config.Unmarshal(&tt)
	if err != nil {
		t.Error(err)
	}
	if !tt.TT.Enable {
		t.Error(`parse failed`)
	}
	if len(tt.ARR) != 2 {
		t.Error(`parse failed`)
	}
	if tt.ARR[0].ITEM != "1" {
		t.Error(`parse failed`)
	}
	if tt.ARR[1].ITEM != "2" {
		t.Error(`parse failed`)
	}
}
