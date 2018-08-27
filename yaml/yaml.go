package yaml

import (
	"io/ioutil"

	db "../db"
	yaml "gopkg.in/yaml.v2"
)

type Configs struct {
	Confs []db.Sshconfig `yaml:"configs"`
}

func ReadYaml(filepath string) Configs {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var d Configs
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	return d
}

func WriteYaml() {
	// TODO
}
