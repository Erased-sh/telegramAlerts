package main

import (
	"fmt"
	"telegramAlerts/internal/realisations"
)

func main() {
	var err error
	bot := realisations.GetEchoBot("")

	client := 23
	action := func() error {
		fmt.Println(client)
		return bot.Bot.SendMsg(0, "Hello, world!")

	}
	bot.Bot.SendMsg(0, "Hello, world!")
	bot.Bot.SendPhoto(0, []byte("photo"))

	if err = bot.SendDelayedMsg(1, "s", action); err != nil {
		fmt.Println(err)
	}
}
