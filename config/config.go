package config

import (
	"os"
	"github.com/elastic/go-elasticsearch/v8"
)

type (
	IConfig interface {
		GetRouteRelativePath() string
		ElasticSearchConfig() (*elasticsearch.TypedClient, error)
	}

	Config struct {}
)

func NewConfig() IConfig {
	return &Config{}
}

func (* Config) GetRouteRelativePath() string {
	return "/api/"
}

func (* Config) ElasticSearchConfig() (*elasticsearch.TypedClient, error){
	host := os.Getenv("ES_HOST")
	config := elasticsearch.Config{
		Addresses: []string{
			host,
		},
	}
	cli, err := elasticsearch.NewTypedClient(config)
	return cli, err
}