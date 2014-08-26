package conf

import (
	"os"
	"io/ioutil"
	yaml "gopkg.in/yaml.v1"
)

type Conf struct {
}

var confInstance *Conf

func init() {
//	err := applyConfigs()
//	if err != nil {
//		panic(err)
//	}
}

func applyConfigs() error {
	confInstance := &Conf{}
	file, err := os.Open("dev.yaml")
	if err != nil {
		return err
	}

		data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, confInstance)
	if err != nil {
		return err
	}
	return nil
}

func Configuration() *Conf {
	return confInstance
}
