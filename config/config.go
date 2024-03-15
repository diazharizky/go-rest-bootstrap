package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Global *viper.Viper

func init() {
	configFilePath := strings.ToLower(os.Getenv("CONFIG_FILE_PATH"))
	if configFilePath == "" {
		configFilePath = "./config"
	}

	configFileName := strings.ToLower(os.Getenv("CONFIG_FILE_NAME"))
	if configFileName == "" {
		configFileName = "default"
	}

	configFileType := strings.ToLower(os.Getenv("CONFIG_FILE_TYPE"))
	if configFileType == "" {
		configFileType = "yaml"
	}

	Global = viper.New()

	Global.AddConfigPath(configFilePath)
	Global.SetConfigName(configFileName)
	Global.SetConfigType(configFileType)
	Global.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	Global.AutomaticEnv()

	if err := Global.ReadInConfig(); err != nil {
		log.Printf("error unable to load config file: %v", err)
	}
}
