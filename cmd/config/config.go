package config

import (
	"encoding/json"
	"log"
	fileHandler "poller-app/internal/service/filehandler"
)

type Config struct {
	AppName        *string `json:"app_name,omitempty"`
	ApiKey         *string `json:"api_key,omitempty"`
	TargetFilePath *string `json:"target_file_path,omitempty"`
	TimeInterval   *int64  `json:"time_interval,omitempty"`
	BaseCurrency   *string `json:"base_currency,omitempty"`
}

func ReadConfig(cp string, logger *log.Logger) (Config, error) {
	confBytes, err := fileHandler.ReadFile(cp, logger)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(confBytes, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
