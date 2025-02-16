package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

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

type tokenSettings struct {
	Key        string `yaml:"SECRET_KEY"`
	MinutesTTL int    `yaml:"TTL_IN_MINUTES"`
}

type userRegSettings struct {
	startBalance int `yaml:"START_BALANCE"`
}

// по-хорошему надо еще наличие элементов внутри файлов конфигов проверять

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

func (c *config) setAuth() error {
	data, err := os.ReadFile("./config/authcfg.yaml")
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}
	a := tokenSettings{}
	err = yaml.Unmarshal(data, &a)
	if err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	c.auth = a

	return nil
}

func (c *config) setUserReg() error {
	data, err := os.ReadFile("./config/userregcfg.yaml")
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}
	u := userRegSettings{}
	err = yaml.Unmarshal(data, &u)
	if err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	c.userReg = u

	return nil
}
