package main

import (
	"os"

	"github.com/mexisme/go-subprocess-mon/subprocess"
	"github.com/mexisme/go-subprocess-mon/version"

	"github.com/mexisme/go-config"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init(config.Config{
		File:       ".subprocess",
		EnvPrefix:  "subprocess",
		Name:       version.Application(),
		Release:    version.Release(),
		FromConfig: true,
	})

	if len(os.Args) < 2 {
		log.Panic("No command provided to execute")
	}
	if err := subprocess.New().SetEnviron(os.Environ()).SetCommand(os.Args[1:]).Run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
