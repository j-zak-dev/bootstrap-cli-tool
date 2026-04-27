package utils

import (
	_ "embed"

	"github.com/spf13/bct/utils/services/builders"
	"github.com/spf13/bct/utils/services/loaders"
	"github.com/spf13/bct/utils/services/writers"
)

func CreateDomain(name string) error {
	var yaml = loaders.LoadDomainYaml()

	var cfg, cfgErr = builders.BuildYamlConfig(yaml)
	if cfgErr != nil {
		panic(cfgErr)
	}

	writers.WriteLayer(name, cfg)

	return nil
}
