package config

// AppConfig allows global access to config
var AppConfig = &AppConfiguration{}

// AppConfiguration is used to define the TOML config structure
type AppConfiguration struct {
	Debug         bool
	ListenAddress string `default:":9055"`

	// XMS Configuration
	XmsEnabled     bool
	XmsV2Enabled   bool
	XmsUser        string `default:"admin"`
	XmsPwd         string `default:"admin"`
	XmsCountersURL string `default:"http://localhost:10080/resource/counters"`
	XmsLicensesURL string `default:"http://localhost:10080/resource/licenses"`
	Xmsv2LicensesURL string `default:"http://localhost:10080/v2/license/stats"`
	Xmsv2CountersURL string `default:"http://localhost:10080/v2/sessions"`

	// C5 Configuration
	SIPProxydEnabled        bool
	SIPProxydExtEnabled     bool
	SIPProxydURL            string `default:"http://127.0.0.1:9980/c5/proxy/commands?49&1&-v"`
	SIPProxydTrunksEnabled  bool
	SIPProxydTrunkStatsURL  string `default:"http://127.0.0.1:9980/c5/proxy/commands?3&7&309"`
	SIPProxydTrunkLimitsURL string `default:"http://127.0.0.1:9980/c5/proxy/commands?3&7&368"`
	SIPProxydSPCountersURL  string `default:"http://127.0.0.1:9980/c5/proxy/commands?4&0&spAll"`
	SIPProxydClSPCountersURL string `default:"http://127.0.0.1:9980/c5/proxy/commands?4&0&spAllCl"`
	ACDQueuedEnabled        bool
	ACDQueuedURL            string `default:"http://127.0.0.1:9982/c5/proxy/commands?49&1&-v"`
	RegistrardEnabled       bool
	RegistrardURL           string `default:"http://127.0.0.1:9984/c5/proxy/commands?49&1&-v"`
	NotificationEnabled     bool
	NotificationURL         string `default:"http://127.0.0.1:9988/c5/proxy/commands?49&1&-v"`
	CstaEnabled             bool
	CstaURL                 string `default:"http://127.0.0.1:9986/c5/proxy/commands?49&1&-v"`

	// Misc
	GoCollectorEnabled      bool
}
