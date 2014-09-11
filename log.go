package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

func newLogger(channellen int64, logPath string) *logs.BeeLogger {
	logger := logs.NewLogger(channellen)

	if *debug {
		logger.SetLogger("console", "")
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", logPath)
		return nil
	}

	return logger
}
