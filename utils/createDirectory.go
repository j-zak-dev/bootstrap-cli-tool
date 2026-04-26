package utils

import (
	"os"
	"strings"
)

func CreateDir(path string) error {
	if strings.Contains(path, "/") {
		return os.MkdirAll(path, 0755)
	} else {
		return os.Mkdir(path, 0755)
	}

}
