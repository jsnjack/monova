package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Cache is an object that stores information about the latest version
type Cache struct {
	Version  *Version
	CommitID string
	path     string
}

// Save Cache object to the cachePath
func (c *Cache) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.path, data, 0644)
}

// Load Cache object
func (c *Cache) load() error {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	return err
}

// CreateCache creates Cache instance
func CreateCache(path string) (*Cache, error) {
	var version Version
	cache := Cache{path: path, Version: &version}
	_, err := os.Stat(path)
	if err != nil {
		if *infoFlag {
			fmt.Println("Cache file doesn't exist. Creating a new file")
		}
		err = cache.Save()
		if err != nil {
			return nil, err
		}
	}
	err = cache.load()
	if err != nil {
		return nil, err
	}
	return &cache, nil
}
