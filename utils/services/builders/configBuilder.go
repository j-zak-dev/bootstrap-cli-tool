package builders

import (
	"gopkg.in/yaml.v2"
)

type LayerConfig struct {
	Layer    string   `yaml:"layer"`
	Elements []string `yaml:"elements"`
}

type AllConfigs map[string]LayerConfig

func BuildYamlConfig(yamlData []byte) (AllConfigs, error) {

	var cfg AllConfigs
	err := yaml.Unmarshal(yamlData, &cfg)

	return cfg, err

}
