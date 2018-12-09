package config

import (
	"os"
	"path/filepath"
)

type AppConfig struct {
	Addr string `default:":8888"`

	SessionID       string `default:"_gream_session"`
	SessionHashKey  string `default:"the-big-and-secret-fash-key-here"`
	SessionBlockKey string `default:"lot-secret-of-characters-big-too"`

	Env string `default:"development"`
}

func (config *AppConfig) IsDevelopment() bool {
	return config.Env == "development"
}

func (config *AppConfig) RootDir() string {
	dir, _ := os.Getwd()
	return dir
}

func (config *AppConfig) Path(paths ...string) string {
	_paths := append([]string{config.RootDir()}, paths...)
	return filepath.Join(_paths...)
}
