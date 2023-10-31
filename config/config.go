package config

import (
	"maunium.net/go/mautrix/bridge/bridgeconfig"
	"maunium.net/go/mautrix/id"
)

type Config struct {
	*bridgeconfig.BaseConfig `yaml:",inline"`

	Analytics struct {
		Host   string `yaml:"host"`
		Token  string `yaml:"token"`
		UserID string `yaml:"user_id"`
	}

	Metrics struct {
		Enabled bool   `yaml:"enabled"`
		Listen  string `yaml:"listen"`
	} `yaml:"metrics"`

	WhatsApp struct {
		OSName      string `yaml:"os_name"`
		BrowserName string `yaml:"browser_name"`
	} `yaml:"whatsapp"`

	Bridge BridgeConfig `yaml:"bridge"`
}

func (config *Config) CanAutoDoublePuppet(userID id.UserID) bool {
	_, homeserver, _ := userID.Parse()
	_, hasSecret := config.Bridge.DoublePuppetConfig.SharedSecretMap[homeserver]
	return hasSecret
}
