package config

import (
	"errors"
	"fmt"
)

var errUninitializedConfig = errors.New("config not initialized")

var cfg config

type config struct {
	initialized bool

	bd      bdParams
	serv    servSettings
	auth    tokenSettings
	userReg userRegSettings
}

// Init необходимая инициализация конфига
func Init() error {
	if err := cfg.setBd(); err != nil {
		return fmt.Errorf("setBd: %w", err)
	}
	if err := cfg.setServ(); err != nil {
		return fmt.Errorf("setServ: %w", err)
	}
	if err := cfg.setAuth(); err != nil {
		return fmt.Errorf("setAuth: %w", err)
	}
	if err := cfg.setUserReg(); err != nil {
		return fmt.Errorf("setUserReg: %w", err)
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
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.bd.Host, c.bd.Port, c.bd.User, c.bd.Password, c.bd.DBname,
	)
}

func (c *config) GetServerAddr() string {
	return c.serv.Addr
}

func (c *config) GetAuthSecretKey() string {
	return c.auth.Key
}

func (c *config) GetAuthTokenTTL() int {
	return c.auth.MinutesTTL
}

func (c *config) GetUserStartBalance() int {
	return c.userReg.startBalance
}
