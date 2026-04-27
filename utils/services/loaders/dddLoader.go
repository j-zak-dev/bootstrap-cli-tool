package loaders

import (
	_ "embed"
)

//go:embed DDD-configs/domain.yaml
var domainYaml []byte

func LoadDomainYaml() []byte {

	return domainYaml

}
