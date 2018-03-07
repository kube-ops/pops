package config

import (
	"os"
	"os/user"

	"github.com/kube-ops/pops/helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitializeConfig Initializes viper with pops presets.
func InitializeConfig() {
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("POPS")
	viper.AutomaticEnv()

	viper.SetConfigName(".pops")

	gitRootDir, err := helper.GitRootDir()
	if err != nil {
		log.Warn(err)

		cwd, err := os.Getwd()
		if err != nil {
			log.Panic(err)
		}

		viper.Set("ProjectRootDir", cwd)
	} else {
		viper.AddConfigPath(gitRootDir)
		viper.Set("ProjectRootDir", gitRootDir)
	}

	usr, userErr := user.Current()
	if userErr != nil {
		log.Panic("Couldn't get the current user.", userErr)
	}

	viper.AddConfigPath(usr.HomeDir)

	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
	}
}
