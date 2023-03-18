package db

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Host     string `json:"host"`
	DBport   int    `json:"dbPort"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

func GetConfiguration() (Config, error) {
	rootDir, err := os.Getwd()
	// To access the config.json file from any location.
	paths := strings.SplitAfter(rootDir, "golang_challenge")
	rootDir = paths[0]

	jsonFile, err := os.Open(filepath.Join(rootDir, "config.json"))
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var config Config

	err = json.Unmarshal(byteValue, &config)

	return config, err
}
