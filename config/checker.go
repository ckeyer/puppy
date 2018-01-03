package config

import (
	"github.com/kubernetes/kubernetes/pkg/volume/validation"
)

func CheckConfig(c *Config) error {
	validation.IsName()
}
