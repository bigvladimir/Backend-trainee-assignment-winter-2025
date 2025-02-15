package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var errUninitializedConfig = errors.New("Config not initialized")

var cfg config

type bdParams struct {
	Host     string `yaml:"POSTGRES_HOST"`
	Port     string `yaml:"POSTGRES_PORT"`
	DBname   string `yaml:"POSTGRES_DB"`
	User     string `yaml:"POSTGRES_USER"`
	Password string `yaml:"POSTGRES_PASSWORD"`
}

type servSettings struct {
	Addr string `yaml:"SERVER_ADDR"`
}

type config struct {
	initialized bool

	bd            bdParams
	serv          servSettings
	authSecretKey string
}

// Init необходимая инициализация конфига
func Init() error {
	if err := cfg.setBd(); err != nil {
		return fmt.Errorf("setBd: %w", err)
	}
	if err := cfg.setServ(); err != nil {
		return fmt.Errorf("setServ: %w", err)
	}
	if err := cfg.setSecretKey(); err != nil {
		return fmt.Errorf("setSecretKey: %w", err)
	}

	cfg.initialized = true

	return nil
}

func Cfg() *config {
	if !cfg.initialized {
		panic(errUninitializedConfig)
	}
	return &cfg
}

func (c *config) GetBdDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.bd.Host, c.bd.Port, c.bd.User, c.bd.Password, c.bd.DBname,
	)
}

func (c *config) GetServerAddr() string {
	return c.serv.Addr
}

func (c *config) GetAuthSecretKey() string {
	return c.authSecretKey
}

func (c *config) setBd() error {
	data, err := os.ReadFile("./config/bdcfg.yaml")
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}
	bd := bdParams{}
	err = yaml.Unmarshal(data, &bd)
	if err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	c.bd = bd

	return nil
}

func (c *config) setServ() error {
	data, err := os.ReadFile("./config/servcfg.yaml")
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}
	serv := servSettings{}
	err = yaml.Unmarshal(data, &serv)
	if err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	c.serv = serv

	return nil
}

func (c *config) setSecretKey() error {
	data, err := os.ReadFile("./config/secretkey.yaml")
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}
	k := struct {
		Key string `yaml:"SECRET_KEY"`
	}{}
	err = yaml.Unmarshal(data, &k)
	if err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	c.authSecretKey = k.Key

	return nil
}
