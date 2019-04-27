package config

import (
	"github.com/jinzhu/configor"
	"time"
)

var Config = &config{}

type config struct {
	AppName                          string

	Mysql struct {
		Host      string
		User      string
		Password  string
		Port      string
		Database  string
		MaxActive int
		MaxIdle   int
		Logdebug  bool
	}

	Redis struct {
		Host        string
		Password    string
		Port        string
		MaxActive   int
		MaxIdle     int
		Idletimeout time.Duration
		Database    int
	}
}

func (c *config) Start(fileName string) {
	err := configor.Load(c, fileName)
	if err != nil {
		panic(err)
	}
}
