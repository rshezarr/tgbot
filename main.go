package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"tgbot/cmd/bot"
)

type Config struct {
	Env      string
	BotToken string
}

func main() {
	configPath := flag.String("config", "config/config.toml", "Path to config file")
	flag.Parse()

	cfg := &Config{}

	_, err := toml.DecodeFile(*configPath, cfg)
	if err != nil {
		log.Fatalf("Error occured while decoding config file: %v", err)
	}

	upgradeBot := bot.UpgradeBot{
		Bot: bot.InitBot(cfg.BotToken),
	}

	upgradeBot.Bot.Handle("/start", bot.StartBotHandler)

	upgradeBot.Bot.Start()
}
