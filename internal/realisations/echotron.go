package realisations

import (
	"context"
	"fmt"
	"github.com/NicoNex/echotron/v3"
	"sync"
)

var (
	echoBot     echoRealisations
	initializer sync.Once
)

type echoRealisations struct {
	ctx context.Context
	echotron.API
	cancel context.CancelFunc
}

func initEchoBot(token string) echoRealisations {
	initializer.Do(func() {
		ctx, cancel := context.WithCancel(context.TODO())
		echoBot = echoRealisations{
			ctx,
			echotron.NewAPI(token),
			cancel}
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
