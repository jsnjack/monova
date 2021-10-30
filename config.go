package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	configFile      = ".monova.config"
	defaultMajorKey = ":M:"
	defaultMinorKey = ":m:"
	defaultPatchKey = ":p:"
)

// Config represents application configuration
type Config struct {
	MajorKeys []string
	MinorKeys []string
	PatchKeys []string
	path      string
}

// Load config file
func (c *Config) Load() error {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	return err
}

// Save config file
func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.path, data, 0644)
}

// CreateConfig creates config instance
func CreateConfig() *Config {
	config := Config{
		MajorKeys: []string{defaultMajorKey},
		MinorKeys: []string{defaultMinorKey},
		PatchKeys: []string{defaultPatchKey},
		path:      configFile,
	}
	_, err := os.Stat(config.path)
	if err != nil {
		// Config does not exist
		config.Save()
		return &config
	}
	err = config.Load()
	if err != nil {
		fmt.Printf("Failed to load config file: %s", err.Error())
	}
	return &config
}
