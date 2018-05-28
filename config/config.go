package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tkanos/gonfig"
)

// LoadConfiguration loads from json file to Configuration
func LoadConfiguration(moduleName string, configuration interface{}) error {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	var configurationFile strings.Builder
	fmt.Fprintf(&configurationFile, "%s/config/%s.json", filepath.Dir(ex), moduleName)
	err = gonfig.GetConf(configurationFile.String(), configuration)
	return err
}
