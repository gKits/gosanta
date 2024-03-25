package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	SMTPHost string `yaml:"smtp_host"`
	SMTPPort string `yaml:"smtp_port"`
	SMTPUser string `yaml:"smtp_user"`
	SMTPPass string `yaml:"smtp_pass"`
}

var configDirs = []string{
	"$XDG_CONFIG_HOME/gosanta/config.yml",
	"$XDG_CONFIG_HOME/gosanta/config.yaml",
	"$XDG_CONFIG_HOME/gosanta.yml",
	"$XDG_CONFIG_HOME/gosanta.yaml",
}

func LoadCfg() (*Cfg, error) {
	cfg := new(Cfg)

	cfg.readEnv()
	if err := cfg.readFromConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Cfg) readFromConfig() error {
	for _, path := range configDirs {
		b, err := os.ReadFile(path)
		if errors.Is(err, os.ErrNotExist) {
			continue
		} else if err != nil {
			return fmt.Errorf("failed to read config file: %s", err.Error())
		}

		if err := yaml.Unmarshal(b, cfg); err != nil {
			return fmt.Errorf("config file at %s is malformed: %s", path, err.Error())
		}
	}

	return nil
}

func (cfg *Cfg) readEnv() {
	cfg.SMTPHost = os.Getenv("GOSANTA_SMTP_HOST")
	cfg.SMTPPort = os.Getenv("GOSANTA_SMTP_PORT")
	cfg.SMTPUser = os.Getenv("GOSANTA_SMTP_USER")
	cfg.SMTPPass = os.Getenv("GOSANTA_SMTP_PASS")
}
