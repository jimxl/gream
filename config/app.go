package config

type AppConfig struct {
	Addr string `default:":8888"`

	SessionID       string `default:"_gream_session"`
	SessionHashKey  string `default:"the-big-and-secret-fash-key-here"`
	SessionBlockKey string `default:"lot-secret-of-characters-big-too"`
}
