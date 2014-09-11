package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var (
	flagSet = flag.NewFlagSet("simages", flag.ExitOnError)

	config  = flagSet.String("config", "", "the path of configure file.(required), default current directory config.ini")
	debug   = flagSet.Bool("debug", false, "run with debug")
	version = flagSet.Bool("version", false, "print version string")
)

func main() {
	flagSet.Parse(os.Args[1:])
	if *version {
		println("0.1")
		os.Exit(0)
	}

	if *config == "" {
		*config = path.Join(currentDir, "config.ini")
	}
	cfg, err := newConfig(*config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	logger := newLogger(10000, cfg.LogPath)
	log.Println(*debug)

	logger.Info("nihao")
	time.Sleep(time.Second * 2)
}
