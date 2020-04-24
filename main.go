package main

import (
	"context"
	"github.com/mail-ru-im/bot-golang"
	"log"
	"os"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if len(botToken) == 0 {
		panic("Bot token not set!")
	}

	bot, err := botgolang.NewBot(botToken)
	if err != nil {
		log.Println("wrong token")
	}

	ctx, _ := context.WithCancel(context.Background())
	updates := bot.GetUpdatesChannel(ctx)
	for update := range updates {
		// your awesome logic here
		chatID := update.Payload.Chat.ID
		message := bot.NewTextMessage(chatID, "Hi, "+update.Payload.From.FirstName)
		err = message.Send()
		if err != nil {
			log.Println(err.Error())
		}
	}

}
