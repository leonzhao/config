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
func New() *MetaConfig {
	var config MetaConfig
	env := os.Getenv("GOENV")
	if env == "" {
		config.Environment = EnvDevelopment
	}

	currentDir, _ := os.Getwd()
	config.Path = filepath.Join(currentDir, DefaultFolder)

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
	extFiles, err := getConfigFiles(mc.Path, mc.Environment)
	for _, v := range extFiles {
		err = parseFile(config, v)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", config)
	}
	return nil
}

// getConfigFiles
// TODO: 可增加读取对应环境变量的文件夹里的所有配置文件名称
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

func parseFile(config interface{}, file string) error {
	fmt.Printf("parse file: %s\n", file)
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
