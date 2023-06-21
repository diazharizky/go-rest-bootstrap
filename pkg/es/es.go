package es

import (
	"github.com/diazharizky/go-rest-bootstrap/config"

	"github.com/elastic/go-elasticsearch/v8"
)

func init() {
	config.Global.SetDefault("elasticsearch.username", "gofiber")
	config.Global.SetDefault("elasticsearch.password", "gofiber")
	config.Global.SetDefault("elasticsearch.addresses", []string{"http://localhost:9200"})
}

func GetClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Username:  config.Global.GetString("elasticsearch.username"),
		Password:  config.Global.GetString("elasticsearch.password"),
		Addresses: config.Global.GetStringSlice("elasticsearch.addresses"),
	}

	return elasticsearch.NewClient(cfg)
}
