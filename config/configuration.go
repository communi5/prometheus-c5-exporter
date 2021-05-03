package config

// AppConfig allows global access to config
var AppConfig = &AppConfiguration{}

// AppConfiguration is used to define the TOML config structure
type AppConfiguration struct {
	Debug                   bool
	ListenAddress           string `default:":9055"`
	JsonDebugEnabled	bool
	XmsDebugEnabled		bool
	// XMS Configuration
	ResourceCountersEnabled 	bool
	ResourceCountersURL 		string `default:"http://localhost:10080/resource/counters"`
	ResourceLicensesEnabled 	bool
	ResourceLicensesURL 		string `default:"http://localhost:10080/resource/licenses"`
	// C5 Configuration
	SIPProxydEnabled        bool
	SIPProxydURL            string `default:"http://127.0.0.1:9980/c5/proxy/commands?49&1&-v"`
	SIPProxydTrunksEnabled  bool
	SIPProxydTrunkStatsURL  string `default:"http://127.0.0.1:9980/c5/proxy/commands?3&7&309"`
	SIPProxydTrunkLimitsURL string `default:"http://127.0.0.1:9980/c5/proxy/commands?3&7&368"`
	ACDQueuedEnabled        bool
	ACDQueuedURL            string `default:"http://127.0.0.1:9982/c5/proxy/commands?49&1&-v"`
	RegistrardEnabled       bool
	RegistrardURL           string `default:"http://127.0.0.1:9984/c5/proxy/commands?49&1&-v"`
    NotificationEnabled     bool
    NotificationURL         string `default:"http://127.0.0.1:9988/c5/proxy/commands?49&1&-v"`
}
