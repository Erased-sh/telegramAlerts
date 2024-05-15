package realisations

import "telegramAlerts/pkg/tick"

func GetEchoBot(token string) tgManager {
	return tgManager{initEchoBot(token)}
}

type tgManager struct {
	Bot TgRealization
}

type TgRealization interface {
	SendMsg(chatId int64, msg string) error
	SendPhoto(chatId int64, file []byte) error
}

type TgWriterDelayed interface {
	SendDelayedMsg(time int64, tag string, action func() error) error
}

func (t tgManager) SendDelayedMsg(time int64, tag string, action func() error) error {
	return tick.GenerateTicker(time, tag, action)
}
