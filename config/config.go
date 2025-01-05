package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port           int    `yaml:"port"`
	EncryptionMode string `yaml:"encryption_mode"`
	LogRetention   int    `yaml:"log_retention_days"`
	StoragePath    string `yaml:"storage_path"`
}

func Load() (*Config, error) {
	configDir := getConfigDir()
	configPath := filepath.Join(configDir, "config.yml")

	// Create default config if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return createDefaultConfig(configPath)
	}

	// Read existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "./"
	}
	return filepath.Join(homeDir, ".config", "climinal")
}

func createDefaultConfig(path string) (*Config, error) {
	config := &Config{
		Port:           8080,
		EncryptionMode: "e2e",
		LogRetention:   30,
		StoragePath:    "./storage",
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(path)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	// Marshal and save default config
	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return nil, err
	}

	return config, nil
}
