package main

import (
	"euterpe/lib/foundation"
	"euterpe/service.discord/bot"
)

func main() {
	f := foundation.New()

	bot.New()
	f.Run()
}
