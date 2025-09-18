package core

import (
	"blog/config"
	"blog/global"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func InitConf() {

	const configFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("get yamlConf  fatal %s", err)
	}
	err1 := yaml.Unmarshal(yamlConf, c)
	if err1 != nil {
		fmt.Printf("config init fatal %s", err1)
	}
	log.Println("config load init success")
	//将配置参数传给glob
	global.Config = c
}
