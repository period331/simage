package main

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	LogPath  string
	Host     string
	HttpPort int
	ScgiPort int
	FontPath string
}

func newConfig(cfg string) (*Config, error) {
	cf, err := goconfig.LoadConfigFile(cfg)

	if err != nil {
		return nil, err
	}

	ds := goconfig.DEFAULT_SECTION

	return &Config{
		LogPath:  cf.MustValue(ds, "logPath", "/tmp/logPath"),
		Host:     cf.MustValue(ds, "bindHost", "0.0.0.0"),
		HttpPort: cf.MustInt(ds, "httpPort", 4230),
		ScgiPort: cf.MustInt(ds, "scgiPort", 4231),
		FontPath: cf.MustValue(ds, "fontPath", currentDir),
	}, nil
}

var currentDir = func() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}()
