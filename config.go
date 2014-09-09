package main

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	configFile *goconfig.ConfigFile
	LogPath    string
	Host       string
	HttpPort   int
	ScgiPort   int
}

func newConfig(cfg string) (*Config, error) {
	config := new(Config)
	configFile, err := goconfig.LoadConfigFile(cfg)
	if err != nil {
		return nil, err
	}
	config.configFile = configFile
	defValConfig := config.getConfigFunc(goconfig.DEFAULT_SECTION)
	defIntConfig := config.getIntConfigFunc(goconfig.DEFAULT_SECTION)
	config.LogPath = defValConfig("logPath", "/tmp/logPath")
	config.Host = defValConfig("bindHost", "0.0.0.0")
	config.HttpPort = defIntConfig("httpPort", 4230)

	config.ScgiPort = defIntConfig("scgiPort", 4231)

	return config, nil
}

func CurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func (c *Config) getConfigFunc(section string) func(key string, def ...string) string {
	return func(key string, def ...string) string {
		return c.configFile.MustValue(goconfig.DEFAULT_SECTION, key, def...)
	}
}

func (c *Config) getIntConfigFunc(section string) func(key string, def ...int) int {
	return func(key string, def ...int) int {
		return c.configFile.MustInt(section, key, def...)
	}
}
