package config

import (
	"checkalive/checker"
	"checkalive/recipient"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Name string `json:"name"`
	Recipient Recipient `json:"recipient"`
	Monitor Monitor `json:"monitor"`
	Template string
}

type Recipient struct {
	Dingtalk [] recipient.DingtalkGroupConfig `json:dingtalk`
}

type Monitor struct {
	Url []checker.UrlCheckRules `json:url`

}

func (config *Config) Load()  {
	configFilePath:=flag.String("c","","json config file")
	flag.Parse()
	jsonFile, err := os.Open(*configFilePath)
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err!= nil {
		fmt.Println("error reading json file")
		return
	}
	err = json.Unmarshal(jsonData, config)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}