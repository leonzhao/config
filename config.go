package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrConfigPath = errors.New("error config path")
)

const (
	DefaultFolder   = "/config"
	DefaultFileName = "default.toml"
)

const (
	EnvDevelopment = "development"
	EnvTesting     = "testing"
	EnvProduction  = "production"
)

// MetaConfig
type MetaConfig struct {
	Environment string
	Path        string
	Verbose     bool
}

// New one instance of MetaConfig
func New(path string) *MetaConfig {
	var config MetaConfig
	env := os.Getenv("GOENV")
	if env == "" {
		config.Environment = EnvDevelopment
	}else {
		config.Environment = env
	}

	if path != "" {
		config.Path = path
	}else {
		currentDir, _ := os.Getwd()
		configPath := filepath.Join(currentDir, "/config")
		config.Path = configPath
	}

	switch config.Environment {
	case EnvDevelopment:
		config.Verbose = true
	case EnvTesting:
		config.Verbose = true
	case EnvProduction:
		config.Verbose = false
	}
	return &config
}

func (mc *MetaConfig) Load(config interface{}) error {
	defer func() {
		if mc.Verbose == true {
			fmt.Printf("loaded config:\n %+v\n", config)
		}
	}()

	fmt.Println("GOENV=", mc.Environment)

	extFiles, err := getConfigFiles(mc.Path, mc.Environment)
	for _, v := range extFiles {
		err = parseFile(config, v)
		if err != nil {
			return err
		}
	}
	return nil
}
