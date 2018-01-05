package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type NetAddr struct {
	Host	string
	Port	int
}

func Read(configFile string, conf interface{}) error {
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
		return err
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		fmt.Printf("read config file %s error:%v", configFile, err)
		return err
	}

	return nil
}
