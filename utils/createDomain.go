package utils

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//go:embed configs/DDD-configs/domain.yaml
var domainYaml []byte

//go:embed configs/DDD-configs/entrypoints.yaml
var entrypointsYaml []byte

//go:embed configs/DDD-configs/infrastructure.yaml
var infrastructureYaml []byte

//go:embed configs/DDD-configs/application.yaml
var applicationYaml []byte

type DomainConfig struct {
	DomainLayer string   `yaml:"domainLayer"`
	Elements    []string `yaml:"elements"`
}

type EntrypointsConfig struct {
	EntrypointLayer string   `yaml:"entrypointLayer"`
	Elements        []string `yaml:"elements"`
}

type InfrastructureConfig struct {
	InfrastructureLayer string   `yaml:"infrastructureLayer"`
	Elements            []string `yaml:"elements"`
}

type ApplicationConfig struct {
	ApplicationLayer string   `yaml:"applicationLayer"`
	Elements         []string `yaml:"elements"`
}

func CreateDomain(name string) error {

	var dmcfg DomainConfig

	dmerr := yaml.Unmarshal(domainYaml, &dmcfg)
	if dmerr != nil {
		panic(dmerr)
	}

	for i := 0; i < len(dmcfg.Elements); i++ {

		path := filepath.Join(name, dmcfg.DomainLayer, dmcfg.Elements[i])

		serviceError := os.MkdirAll(path, 0755)
		if serviceError != nil {
			fmt.Println("Error creating layer: " + path)
			return serviceError
		}

	}

	var encfg EntrypointsConfig

	enerr := yaml.Unmarshal(entrypointsYaml, &encfg)
	if enerr != nil {
		panic(enerr)
	}

	for i := 0; i < len(encfg.Elements); i++ {

		path := filepath.Join(name, encfg.EntrypointLayer, encfg.Elements[i])

		serviceError := os.MkdirAll(path, 0755)
		if serviceError != nil {
			fmt.Println("Error creating layer: " + path)
			return serviceError
		}

	}

	var incfg InfrastructureConfig

	inerr := yaml.Unmarshal(infrastructureYaml, &incfg)
	if inerr != nil {
		panic(inerr)
	}

	for i := 0; i < len(incfg.Elements); i++ {

		path := filepath.Join(name, incfg.InfrastructureLayer, incfg.Elements[i])

		serviceError := os.MkdirAll(path, 0755)
		if serviceError != nil {
			fmt.Println("Error creating layer: " + path)
			return serviceError
		}

	}

	var apcfg ApplicationConfig

	aperr := yaml.Unmarshal(applicationYaml, &apcfg)
	if aperr != nil {
		panic(aperr)
	}

	for i := 0; i < len(apcfg.Elements); i++ {

		path := filepath.Join(name, apcfg.ApplicationLayer, apcfg.Elements[i])

		serviceError := os.MkdirAll(path, 0755)
		if serviceError != nil {
			fmt.Println("Error creating layer: " + path)
			return serviceError
		}

	}
	return nil
}
