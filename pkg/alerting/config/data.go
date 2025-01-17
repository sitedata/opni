package config

import (
	cfg "github.com/prometheus/alertmanager/config"
	"gopkg.in/yaml.v3"
)

// Mimics github.com/prometheus/alertmanager/config/config.go's Config struct
// but we can't due to mismatched github.com/prometheus/common versions
type ConfigMapData struct {
	Global       *GlobalConfig      `yaml:"global,omitempty" json:"global,omitempty"`
	Route        *cfg.Route         `yaml:"route,omitempty" json:"route,omitempty"`
	InhibitRules []*cfg.InhibitRule `yaml:"inhibit_rules,omitempty" json:"inhibit_rules,omitempty"`
	Receivers    []*Receiver        `yaml:"receivers,omitempty" json:"receivers,omitempty"`
	Templates    []string           `yaml:"templates" json:"templates"`
}

func (c *ConfigMapData) Parse(data string) error {
	return yaml.Unmarshal([]byte(data), c)
}

func NewConfigMapDataFrom(data string) (*ConfigMapData, error) {
	c := &ConfigMapData{}
	err := c.Parse(data)
	return c, err
}

func (c *ConfigMapData) Marshal() ([]byte, error) {
	data, err := yaml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return data, nil
}
