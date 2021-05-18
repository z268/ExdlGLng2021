package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Config struct {
	Db struct {
		Driver   string
		Database string
		Host     string
		Port     string
		User     string
		Password string
	}
	Http struct {
		Port uint16
		Host string
	}
	CatalogGrpc struct {
		Host string
		Port uint64
	}
}

func Init(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if err = yaml.Unmarshal(buf, c); err != nil {
		return nil, err
	}

	c.Db.Driver = getEnv("DB_DRIVER", "mysql")
	c.Db.Host = getEnv("DB_HOST", "localhost")
	c.Db.Port = getEnv("DB_PORT", "3306")
	c.Db.User = getEnv("DB_USER", "root")
	c.Db.Password = getEnv("DB_PASSWORD", "task6")
	c.Db.Database = getEnv("DB_DATABASE", "orderdb")

	return c, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

