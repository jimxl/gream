package config

import "github.com/koding/multiconfig"

var App = new(AppConfig)

func init() {
	// TODO: 这个地方应该要导入默认的配置文件
	m := multiconfig.New()
	m.MustLoad(App)
}
