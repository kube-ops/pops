package config

import (
	"os"
	"os/user"
	path "path/filepath"

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

	currDir, err := os.Getwd()
	if err != nil {
		log.Panic("Couldn't initialize configuration.", err)
	}

	for path.Clean(currDir) != "/" {
		isGitRootDir, fileStatErr := helper.Exists(path.Join(currDir, ".git"))

		if fileStatErr != nil {
			log.Warn("Error while searching the project git root", err)
		}

		if isGitRootDir {
			viper.AddConfigPath(currDir)
			break
		}

		currDir, _ = path.Split(currDir)
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
