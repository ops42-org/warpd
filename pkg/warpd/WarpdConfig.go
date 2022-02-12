package warpd

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type WarpdConfig struct {
	Build      []WarpdConfigBuild      `yaml:"build"`
	EnvMapping []WarpdConfigEnvMapping `yaml:"envMapping"`
}

func LoadConfig(configPath string) (*WarpdConfig, error) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: ", err)
	}

	var warpdConfig WarpdConfig

	err = yaml.Unmarshal(yamlFile, &warpdConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal yaml: ", err)
	}

	return &warpdConfig, nil
}

type WarpdConfigBuild struct {
	Path       *string           `yaml:"path"`
	Builder    *string           `yaml:"builder"`
	Buildpacks []string          `yaml:"buildpacks"`
	Env        map[string]string `yaml:"env"`
}

type WarpdConfigEnvMapping struct {
	Branch          *string  `yaml:"branch"`
	EnvName         *string  `yaml:"envName"`
	Cluster         *string  `yaml:"cluster"`
	ExcludeBranches []string `yaml:"excludeBranches"`
}
