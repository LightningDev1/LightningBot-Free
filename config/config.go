package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/LightningDev1/LightningBot-Free/utils"
)

type Config struct {
	Token         string
	CommandPrefix string
	Embed         EmbedConfig
}

type EmbedConfig struct {
	Title  string
	Footer string
}

var (
	lastConfig     *Config
	lastConfigTime time.Time
	configMutex    = &sync.Mutex{}
)

func (c *Config) Save() error {
	configMutex.Lock()
	defer configMutex.Unlock()

	// Open the config file
	configFile, err := os.Create(filepath.Join(utils.File.GetMainFolder(), "Config", "config.json"))
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Encode the config
	configBytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// Write the config to the file
	_, err = configFile.Write(configBytes)
	if err != nil {
		return err
	}

	// Cache the config
	lastConfig = c
	lastConfigTime = time.Now()

	return nil
}

func Load() (config Config, err error) {
	configMutex.Lock()
	defer configMutex.Unlock()

	if isLastConfigValid() {
		return *lastConfig, nil
	}

	// Open the config file
	configFile, err := os.Open(filepath.Join(utils.File.GetMainFolder(), "Config", "config.json"))
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	// Decode the config file
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return config, err
	}

	// Cache the config
	lastConfig = &config
	lastConfigTime = time.Now()

	return config, nil
}

func isLastConfigValid() bool {
	// Check if the config has already loaded in the last 2 seconds
	return lastConfig != nil && time.Since(lastConfigTime) < time.Second*10
}
