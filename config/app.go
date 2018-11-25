package config

type AppConfig struct {
	Addr string `default:":8888"`

	SessionID string `default:"5ab421c4a9bd41fb9a6284159e0b21ebadbe59cc0d5435a7f8a7e9b1bf87a75b7ed87e14a31af8fcb14eeb0dded9d8595a4e2a346b4a54860bcc4cd2c1d246a8"`
}
