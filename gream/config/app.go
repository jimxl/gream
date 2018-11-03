package config

import "time"

type AppConfig struct {
	Addr         string        `default:":8888"`
	WriteTimeout time.Duration `default:(15 * time.Second)`
	ReadTimeout  time.Duration `default:(15 * time.Second)`
}
