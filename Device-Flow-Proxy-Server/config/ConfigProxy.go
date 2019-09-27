package config

import (
"gopkg.in/yaml.v2"
"io/ioutil"
"log"

)

type conf struct {
	URI string `yaml:"verification_uri"`
	Time int64 `yaml:"time"`
	ClientId string `yaml:"client_id"`
	GrantType string `yaml:"grant_type"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func GetVerificationURI() (out string) {
	var c conf
	c.getConf()
	return c.URI
}

func GetClientId() (out string) {
	var c conf
	c.getConf()
	return c.ClientId
}

func GetGrantType() (out string) {
	var c conf
	c.getConf()
	return c.GrantType
}

