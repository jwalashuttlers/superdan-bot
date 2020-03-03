package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

// cfgTelegram contains config regarding telegram
type cfgTelegram struct {
	APIKey string `koanf:"api_key"`
}

// Config represents app config
type Config struct {
	Telegram cfgTelegram
}

var cfg Config

func initConfig() {
	var k = koanf.New(".")

	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Fatalf("couldn't read config: %v", err)
	}

	k.Unmarshal("telegram", &cfg.Telegram)
}
