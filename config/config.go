package config

import (
	"github.com/elastic/go-elasticsearch/v7"
	"os"
)

type (
	IConfig interface {
		GetRouteRelativePath() string
		ElasticSearchConfig() (*elasticsearch.Client, error)
	}

	Config struct {}
)

func NewConfig() IConfig {
	return &Config{}
}

func (* Config) GetRouteRelativePath() string {
	return "/api/"
}

func (* Config) ElasticSearchConfig() (*elasticsearch.Client, error){
	host := os.Getenv("ES_HOST")
	config := elasticsearch.Config{
		Addresses: []string{
			host,
		},
	}
	cli, err := elasticsearch.NewClient(config)
	return cli, err
}