package utils

import (
	"os"
	"path/filepath"
)

type FileUtils struct{}

func (FileUtils) GetMainFolder() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return filepath.Join(configDir, "LightningBotFree")
}

func (FileUtils) CreateDirectories() error {
	mainFolder := File.GetMainFolder()

	directories := []string{
		mainFolder,
		filepath.Join(mainFolder, "Config"),
	}

	for _, directory := range directories {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}
	
	return nil
}

var File = &FileUtils{}
