package realisations

import (
	"context"
	"telegramAlerts/pkg/tick"
)

func GetEchoBot(ctx context.Context, token string) tgManager {
	ctx, cancel := context.WithCancel(ctx)
	return tgManager{
		ctx,
		initEchoBot(token),
		cancel,
	}
}

type TgRealization interface {
	SendMsg(chatId int64, msg string) error
	SendPhoto(chatId int64, file []byte) error
}

type tgManager struct {
	ctx    context.Context
	Bot    TgRealization
	cancel context.CancelFunc
}

func (t tgManager) SendDelayedMsg(time int64, tag string, action func() error) {
	tick.GenerateTicker(t.ctx, time, tag, action)
}

func (t tgManager) CancelAllRealizationBroadcastingMsg() { //cancel all async alerts of exactly one realisation
	t.cancel()
}
