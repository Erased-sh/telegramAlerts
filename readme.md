## Bot service
Библиотека позволяет создавать заранее подготовленные модели часто используемых реализаций telegram, slack и discord ботов, обобщая их основные методы.

В дополнение, telegram bot service учитывает сложившиеся паттерны проектирования ботов в качестве асинхронной системы оповещений, реализуя формат отложенных операций с оповещением.

Архитектура:

Config:
Унифицирует структы в конфигах

Realisation:
В этом разделе представлены существующие библиотеки для взаимодействия с api telegram/slack/discord
Каждая реализация должна интегрировать внутреннюю логику сторонних библиотек в одну единую по методам sendPhoto, sendMsg
-bot.go
создает уникальный объект реализации со своим собственным контекстом
позволяет вызвать методы sendMsg,sendPhot
реализует асинхронную отправку event в чат
Ticker:
-позволяет запустить асинхронный алерт по истечению времени. Логика алерта определяется внутри области видимости вызывающей функции, поэтому сохраняет доступ ко всем переменным

Context:
Может быть отменен глобально для всех реализаций, либо для всех тикеров одного объекта

Sync msg
```golang 
ctx, cancel := context.WithCancel(context.Background())
	bot := realisations.GetEchoBot(ctx, "TOKEN")
    bot.Bot.SendMsg(chatid, "Simple sync msg")
```

Async msg
```golang 
ctx, cancel := context.WithCancel(context.Background())
	bot := realisations.GetEchoBot(ctx, "TOKEN")

	tries := 0
	action := func() error {
		tries++
		return bot.Bot.SendMsg(7138652177, fmt.Sprintf("Async message. Try №%d", tries))
	}
	
	go bot.SendDelayedMsg(1, "s", action)

	time.Sleep(10 * time.Second)
	cancel()
	//unreachable child thread code
	time.Sleep(10 * time.Second)
```

