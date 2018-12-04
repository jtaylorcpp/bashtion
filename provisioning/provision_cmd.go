package provisioning

import (
	"gopkg.in/yaml.v2"
)

type Command struct {
	Kind    string `yaml:"kind"`
	Type    string `yaml:"type"`
	Content string `yaml:"content"`
}

func CommandFromYamlString(command_yaml []byte) Command {
	cmd := &Command{}
	err := yaml.Unmarshal(command_yaml, cmd)
	if err != nil {
		panic(err)
	}

	return *cmd
}
