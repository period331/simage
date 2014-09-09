package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	flagSet = flag.NewFlagSet("simages", flag.ExitOnError)

	config  = flagSet.String("config", "", "the path of configure file.(required), default current directory config.ini")
	version = flagSet.Bool("version", false, "print version string")
)

func main() {
	flagSet.Parse(os.Args[1:])
	if *version {
		println("0.1")
		os.Exit(0)
	}

	if *config == "" {
		*config = path.Join(CurrentDir(), "config.ini")
	}
	cfg, err := newConfig(*config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	fmt.Println(CurrentDir())
}
