package config

import (
	"encoding/json"
	"os"

	"BT_supertoolsV2/devices"
)

type RawConfig struct {
	Devices []devices.DeviceConfig `json:"devices"`
}

func LoadConfig(file string, dst *[]devices.DeviceConfig) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var conf RawConfig
	if err := json.Unmarshal(data, &conf); err != nil {
		return err
	}
	*dst = conf.Devices
	return nil
}
