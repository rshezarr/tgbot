package bot

import (
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type UpgradeBot struct {
	Bot *telebot.Bot
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Error while initializing bot: %v", err)
	}

	return b
}

func StartBotHandler(context telebot.Context) error {
	return context.Send("Hello, " + context.Sender().FirstName)
}
