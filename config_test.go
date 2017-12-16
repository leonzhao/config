package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"os"
)

type Config struct {
	Name string
	Branch string
	DB DBConfig `toml:"database"`
	Redis RedisConfig `toml:"redis"`
}

type RedisConfig struct {
	Host string
	Port uint
}

type DBConfig struct {
	Host      string
	Port      uint
	Username  string
	Password  string
	DBName    string
	Charset   string
	ParseTime bool
}

func TestNew(t *testing.T) {
	mc := New()
	assert.Equal(t, EnvDevelopment, mc.Environment)
	assert.Equal(t, true, mc.Verbose)

	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/config")
	assert.Equal(t, configPath, mc.Path)
}

func TestLoad(t *testing.T) {
	mc := New()

	var config Config
	mc.Load(&config)
}