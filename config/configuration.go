package config

// AppConfig allows global access to config
var AppConfig AppConfiguration

// AppConfiguration is used to define the TOML config structure
type AppConfiguration struct {
	Debug             bool
	ListenAddress     string `default:":9055"`
	SIPProxydEnabled  bool   // `default:"true"`
	SIPProxydURL      string `default:"http://127.0.0.1:9980/c5/proxy/commands?49&1&-v"`
	ACDQueuedEnabled  bool   `default:"true"`
	ACDQueuedURL      string `default:"http://127.0.0.1:9982/c5/proxy/commands?49&1&-v"`
	RegistrardEnabled bool   `default:"true"`
	RegistrardURL     string `default:"http://127.0.0.1:9984/c5/proxy/commands?49&1&-v"`
}
