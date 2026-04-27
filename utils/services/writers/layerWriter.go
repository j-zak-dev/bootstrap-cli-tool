package writers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/bct/utils/services/builders"
)

func WriteLayer(name string, cfg builders.AllConfigs) error {

	for layerKey, layerCfg := range cfg {

		for i := 0; i < len(layerCfg.Elements); i++ {

			path := filepath.Join(name, layerCfg.Layer, layerCfg.Elements[i])

			err := os.MkdirAll(path, 0755)
			if err != nil {
				fmt.Println("Error creating layer:", path)
				return err
			}
		}

		_ = layerKey // optional if not used
	}

	return nil
}
