package main

import (
	"github.com/kube-ops/pops/cmd"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	InitializeConfig()

	if err := ReadConfigFile(); err != nil {
		log.Warn(err)
	}

	cmd.Execute()
}
