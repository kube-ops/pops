package main

import (
	"github.com/spf13/viper"

	"os"
	"os/user"
	path "path/filepath"
)

// InitializeConfig Initializes viper with pops presets.
func InitializeConfig() {
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("POPS")
	viper.AutomaticEnv()

	viper.SetConfigName(".pops")

	currDir, _ := os.Getwd()

	for path.Clean(currDir) != "/" {
		if fileExists, _ := exists(path.Join(currDir, ".git")); fileExists {
			viper.AddConfigPath(currDir)
			break
		}
		currDir, _ = path.Split(currDir)
	}

	usr, _ := user.Current()
	viper.AddConfigPath(usr.HomeDir)

	viper.AddConfigPath("/etc")
}

// ReadConfigFile reads the config file.
func ReadConfigFile() error {
	return viper.ReadInConfig()
}
