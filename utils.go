package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// getConfigFiles get all config files based on environment variable `GOENV`
func getConfigFiles(path string, env string) ([]string, error) {
	var files []string

	if path == "" {
		return nil, ErrConfigPath
	}

	defaultConfig := filepath.Join(path, DefaultFileName)
	files = append(files, defaultConfig)

	defaultEnvConfig := filepath.Join(path, env+".toml")
	files = append(files, defaultEnvConfig)

	extConfigFolder := filepath.Join(path, env)
	err := filepath.Walk(extConfigFolder, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.Mode().IsRegular() {
			files = append(files, path)
			return nil
		}
		return nil
	})
	return files, err
}

// parseFile parse config from file
func parseFile(config interface{}, file string) error {
	fmt.Printf("parse config: %s\n", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
		return yaml.Unmarshal(data, config)
	case strings.HasSuffix(file, ".toml"):
		_, err := toml.DecodeFile(file, config)
		return err
	case strings.HasSuffix(file, ".json"):
		return json.Unmarshal(data, config)
	}
	return errors.New("error file type")
}
