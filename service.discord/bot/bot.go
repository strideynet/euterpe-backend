package bot

import (
	"context"
	"euterpe/lib/log"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"go.uber.org/zap"
	"os"
)

var logger = log.Package("service.discord:bot")

type Bot struct {
	dc *disgord.Client
}

func (b *Bot) Run() {
	prefix := ".euterpe"

	logFilter, _ := std.NewLogFilter(b.dc)
	filter, _ := std.NewMsgFilter(context.Background(), b.dc)
	filter.SetPrefix(prefix)

	b.dc.Gateway().
		WithMiddleware(filter.NotByBot, filter.HasPrefix, logFilter.LogMsg, filter.StripPrefix).
		MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
			// Initiate context for message
			logger.Info(h.Message.Content)
			ctx := context.Background()
			_, _ = h.Message.Reply(ctx, s, "Absolutely poggers")
		})

	err := b.dc.Gateway().StayConnectedUntilInterrupted()
	if err != nil {
		logger.Fatal("Err occurred starting bot", zap.Error(err))
	}
}

func New() *Bot {
	b := &Bot{}
	b.dc = disgord.New(disgord.Config{
		ProjectName: "euterpe",
		BotToken:    os.Getenv("DISGORD_TOKEN"), // TODO: replace this with config loading
		Logger:      log.Package("disgord").Sugar(),
		RejectEvents: []string{
			disgord.EvtTypingStart,

			// These require sepecial intents
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},

		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "https://github.com/strideynet/euterpe-backend",
			},
		},
	})

	return b
}
