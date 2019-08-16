package util

import (
	"io/ioutil"

	"github.com/donh/shortner/pkg/models"
	"gopkg.in/yaml.v2"
)

// Config parses the content of config/config.yml
func Config() models.ConfigStruct {
	config := models.ConfigStruct{}
	source, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		source, err = ioutil.ReadFile("../config/config.yml")
		if err != nil {
			source, err = ioutil.ReadFile("../../config/config.yml")
			if err != nil {
				panic(err)
			}
		}
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	return config
}
