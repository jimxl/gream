package config

type AppConfig struct {
	Port string `default:"8888"`
	Host string `default:""`
}
