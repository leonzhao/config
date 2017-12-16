package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

type Config struct {
	Name   string
	Branch string
	DB     DBConfig    `toml:"database"`
	Redis  RedisConfig `toml:"redis"`
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
	assert.Equal(t, "metaconfig", config.Name)
	assert.Equal(t, "dev", config.Branch)
	assert.Equal(t, "localhost", config.DB.Host)
	assert.Equal(t, uint(3500), config.DB.Port)
	assert.Equal(t, "root", config.DB.Username)
	assert.Equal(t, "mc_dev", config.DB.Password)
	assert.Equal(t, "config", config.DB.DBName)
	assert.Equal(t, "utf8", config.DB.Charset)
	assert.Equal(t, true, config.DB.ParseTime)
	assert.Equal(t, "127.0.0.1", config.Redis.Host)
	assert.Equal(t, uint(4000), config.Redis.Port)
}
