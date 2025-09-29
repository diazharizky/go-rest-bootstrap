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
		configFileName = "base"
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

	setAppDefaultConfig()
	setDbDefaultConfig()
	setEmailClientDefaultConfig()
}

func setAppDefaultConfig() {
	Global.SetDefault("app.host", "localhost")
	Global.SetDefault("app.port", 8080)
	Global.SetDefault("app.throttling.max.requests", 20)
	Global.SetDefault("app.throttling.expiration", 30)
	Global.SetDefault("app.jwt_secret", os.Getenv("APP_JWT_SECRET"))
}

func setDbDefaultConfig() {
	Global.SetDefault("db.host", "localhost")
	Global.SetDefault("db.port", 5432)
	Global.SetDefault("db.user", "gorestbs")
	Global.SetDefault("db.password", "gorestbs")
	Global.SetDefault("db.name", "gorestbs")
}

func setEmailClientDefaultConfig() {
	Global.SetDefault("emailclient.host", "localhost")
	Global.SetDefault("emailclient.port", 1025)
	Global.SetDefault("emailclient.sender_name", "gorest")
	Global.SetDefault("emailclient.email", "gorest")
	Global.SetDefault("emailclient.password", "")
}
