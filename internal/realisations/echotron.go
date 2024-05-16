package realisations

import (
	"fmt"
	"github.com/NicoNex/echotron/v3"
	"sync"
)

var (
	echoBot     echoRealisations
	initializer sync.Once
)

type echoRealisations struct {
	echotron.API
}

func initEchoBot(token string) echoRealisations {
	initializer.Do(func() {
		echoBot = echoRealisations{echotron.NewAPI(token)}
	})
	return echoBot
}

func (e echoRealisations) SendMsg(chatId int64, msg string) error {
	if msg == "" {
		return fmt.Errorf("tg bot cannot send empty msg")
	}
	_, err := e.API.SendMessage(msg, chatId, nil)
	return err
}

func (e echoRealisations) SendPhoto(chatId int64, file []byte) error {
	if file == nil {
		return fmt.Errorf("tg bot cannot send empty file")
	}
	photo := echotron.InputFile{}
	_, err := e.API.SendPhoto(photo, chatId, nil)
	return err
}
