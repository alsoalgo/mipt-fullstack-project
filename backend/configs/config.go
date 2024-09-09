package configs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type DB struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
}

type Config struct {
	DB DB
}

func Parse(env string) (*Config, error) {
	dbConfig, err := parseDBConfig(env)
	if err != nil {
		return nil, err
	}
	return &Config{
		DB: *dbConfig,
	}, nil
}

func parseDBConfig(env string) (*DB, error) {
	filename, err := filepath.Abs(fmt.Sprintf("./configs/%s/db.yml", env))
	if err != nil {
		return nil, errors.New("Can't find db config file:" + err.Error())
	}
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Can't read db config file:" + err.Error())
	}
	var dbConfig *DB
	err = yaml.Unmarshal(yamlFile, &dbConfig)
	if err != nil {
		return nil, errors.New("Can't parse db config file:" + err.Error())
	}
	return dbConfig, nil
}
