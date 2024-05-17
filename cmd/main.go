package main

import (
	"context"
	"fmt"
	"telegramAlerts/internal/realisations"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	bot := realisations.GetEchoBot(ctx, "7109287437:AAFpz57MsiPG1vsUEAYhX1N20rHtZChhBMA")

	tries := 0
	action := func() error {
		tries++
		return bot.Bot.SendMsg(7138652177, fmt.Sprintf("Async message. Try №%d", tries))
	}

	bot.Bot.SendMsg(7138652177, "Simple sync msg")
	go bot.SendDelayedMsg(1, "s", action)

	time.Sleep(10 * time.Second)
	cancel() //U can cancel broadcastiong of all realisations
	time.Sleep(10 * time.Second)
}
