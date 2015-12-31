package config

import (
	"os"
	"os/user"

	"github.com/spf13/viper"

	"github.com/highrisehq/github-issues/logger"
)

const (
	fileName = ".github-issues"
	fileType = "toml"
)

var filePath string
var ConfigFile string

func init() {
	usr, _ := user.Current()
	filePath = usr.HomeDir
}

func GetToken() string {
	return viper.GetString("authentication.token")
}

// Read in config file and ENV variables if set.
func ConfigInit() {
	viper.SetConfigType(fileType)
	if ConfigFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(ConfigFile)
	}

	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.AddConfigPath(filePath) // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Debug("Using config file: " + viper.ConfigFileUsed())
	}
}

func GenerateConfig() {
	if _, err := os.Stat(fullFileLocation()); err == nil {
		logger.Debug("Config file already exists")
	} else {
		logger.Debug("Generating new config file at " + fullFileLocation() + "...")
	}
}

func fullFileLocation() string {
	return filePath + "/" + fileName + "." + fileType
}
