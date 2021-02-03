package main

import (
	"euterpe/lib/foundation"
	"euterpe/service.discord/bot"
)

func main() {
	f := foundation.New()

	b := bot.New()

	// TODO: Proper exiting handling
	go func() {
		b.Run()
	}()
	f.Run()
}
