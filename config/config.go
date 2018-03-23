package config

import (
	"os"
	"os/user"
	"path"

	"github.com/kube-ops/pops/helper"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var configName = ".pops"
var configType = "yaml"

// InitializeConfig Initializes viper with pops presets.
func InitializeConfig() {
	viper.SetConfigType(configType)
	viper.SetEnvPrefix("POPS")
	viper.AutomaticEnv()

	viper.SetConfigName(configName)

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

// SafeWriteConfig write the configuration file if not exists.
func SafeWriteConfig() error {
	configFile := path.Join(viper.GetString("ProjectRootDir"), configName+"."+configType)

	exists, err := helper.Exists(configFile)

	if err != nil {
		return err
	}

	if exists {
		log.Info("Configuration file already exists. Skipping!!! Path: ", configFile)
		return nil
	}

	// Necessary because viper write function returns an error when file does not exists
	if _, err := os.OpenFile(configFile, os.O_RDONLY|os.O_CREATE, 0600); err != nil {
		log.Error("Failed to create configuration file: ", configFile)
		return err
	}

	log.Info("Creating default configuration file. => ", configFile)
	return viper.WriteConfigAs(configFile)
}
