package bot

import (
	"euterpe/lib/log"
	"github.com/andersfylling/disgord"
	"os"
)

type Bot struct {
	dc *disgord.Client
}

func New() *Bot {
	b := &Bot{}
	b.dc = disgord.New(disgord.Config{
		BotToken: os.Getenv("DISGORD_TOKEN"), // TODO: replace this with config loading
		Logger:   log.Package("disgord").Sugar(),
	})

	return b
}
