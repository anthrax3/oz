package oz

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ProfileDir      string   `json:"profile_dir"`
	ShellPath       string   `json:"shell_path"`
	SandboxPath     string   `json:"sandbox_path"`
	BridgeMACAddr   string   `json:"bridge_mac"`
	NMIgnoreFile    string   `json:"nm_ignore_file"`
	UseFullDev      bool     `json:"use_full_dev"`
	AllowRootShell  bool     `json:"allow_root_shell"`
	LogXpra         bool     `json:"log_xpra"`
	EnvironmentVars []string `json:"environment_vars"`
}

const DefaultConfigPath = "/etc/oz/oz.conf"

func NewDefaultConfig() *Config {
	return &Config{
		ProfileDir:      "/var/lib/oz/cells.d",
		ShellPath:       "/bin/bash",
		SandboxPath:     "/srv/oz",
		NMIgnoreFile:    "/etc/NetworkManager/conf.d/oz.conf",
		BridgeMACAddr:   "6A:A8:2E:56:E8:9C",
		UseFullDev:      false,
		AllowRootShell:  false,
		LogXpra:         false,
		EnvironmentVars: []string{
			"USER", "USERNAME", "LOGNAME",
			"LANG", "LANGUAGE", "_",
		},
	}
}

func LoadConfig(path string) (*Config, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := NewDefaultConfig()
	if err := json.Unmarshal(bs, c); err != nil {
		return nil, err
	}
	return c, nil
}
