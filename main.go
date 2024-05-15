package main

import (
	"fmt"
	"telegramAlerts/internal/realisations"
)

func main() {
	var err error
	bot := realisations.GetEchoBot("")
	client := 23
	if err = bot.SendDelayedMsg(1, "s", func() error {
		fmt.Println(client)
		return bot.Bot.SendMsg(0, "Hello, world!")
	}); err != nil {
		fmt.Println(err)
	}
}
