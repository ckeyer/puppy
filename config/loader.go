package config

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/ckeyer/logrus"
	yaml "gopkg.in/yaml.v2"
)

func LoadFile(name string) (*Config, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("open config file %s successful", name)
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	return Load(buf.Bytes())
}

func Load(data []byte) (cfg *Config, err error) {
	cfgData := bytes.TrimSpace(data)
	if bytes.HasPrefix(cfgData, []byte("{")) {
		cfg, err = loadJSON(cfgData)
	} else {
		cfg, err = loadYAML(cfgData)
	}
	if err != nil {
		return nil, err
	}

	logrus.Infof("load config successful.")
	return cfg, nil
}

func loadYAML(in []byte) (*Config, error) {
	cfg := &Config{}
	err := yaml.Unmarshal(in, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func loadJSON(in []byte) (*Config, error) {
	cfg := &Config{}
	err := json.Unmarshal(in, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
