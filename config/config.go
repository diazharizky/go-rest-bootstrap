package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Global *viper.Viper

func init() {
	Global = viper.New()

	configFilePath := strings.ToLower(os.Getenv("CONFIG_FILE_PATH"))
	if configFilePath == "" {
		configFilePath = "./config"
	}
	Global.AddConfigPath(configFilePath)

	configFileName := strings.ToLower(os.Getenv("CONFIG_FILE_NAME"))
	if configFileName == "" {
		configFileName = "default"
	}
	Global.SetConfigName(configFileName)

	configFileType := strings.ToLower(os.Getenv("CONFIG_FILE_TYPE"))
	if configFileType == "" {
		configFileType = "yaml"
	}
	Global.SetConfigType(configFileType)

	if err := Global.ReadInConfig(); err != nil {
		log.Printf("error unable to load config file: %v", err)
	}

	Global.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Global.AutomaticEnv()
}
