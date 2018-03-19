package nsshconfig

import (
	"errors"
	"os"
)

func WriteConfig() error {
	var config []string
	var content string

	tmpName := TempFileName("config_", ".nssh")

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(tmpName, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return err
	}

	defer f.Close()

	for _, x := range Hosts() {
		config = append(config, x.Decode())
	}

	if len(config) == 0 {
		return errors.New("Not entries to write")
	}

	for _, c := range config {
		content += c

	}

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	err = os.Rename(tmpName, *configPath)
	if err != nil {
		return err
	}
	return nil
}