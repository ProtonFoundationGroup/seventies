package main

import (
	"github.com/ProtonFundationGroup/seventies/v2/config"
	"github.com/ProtonFundationGroup/seventies/v2/log"
	"gopkg.in/yaml.v2"
)

func main() {
	configFile := "../seventies.yml"
	conf, err := config.Load(configFile)
	if err != nil {
		log.PrintNormal("load config file `%v` failed, err = %v", configFile, err)
		return
	}

	configStr, err := yaml.Marshal(conf)
	if err != nil {
		return
	}

	log.PrintGreen("config => \n%s", configStr)
}
