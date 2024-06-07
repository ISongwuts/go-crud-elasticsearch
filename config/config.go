package config

type (
	IConfig interface {
		GetRouteRelativePath() string
	}

	Config struct {}
)

func NewConfig() IConfig {
	return &Config{}
}

func (* Config) GetRouteRelativePath() string {
	return "/api/"
}