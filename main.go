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
		log.Fatalf("Bot token not set!")
	}

	bot, err := botgolang.NewBot(botToken, botgolang.BotDebug(true))
	if err != nil {
		log.Fatalf("Сannot connect to bot: %s", err)
	}

	log.Println(bot.Info)

	ctx, _ := context.WithCancel(context.Background())
	updates := bot.GetUpdatesChannel(ctx)
	for update := range updates {

		log.Println(update.Type, update.Payload)
		switch update.Type {
		case botgolang.NEW_MESSAGE:
			message := update.Payload.Message()

			helloBtn := botgolang.NewCallbackButton("Hello", "echo")
			goBtn := botgolang.NewURLButton("go", "https://golang.org/")
			message.AttachInlineKeyboard([][]botgolang.Button{{helloBtn, goBtn}})

			if err := message.Send(); err != nil {
				log.Printf("failed to send message: %s", err)
			}
		case botgolang.EDITED_MESSAGE:
			message := update.Payload.Message()
			if err := message.Reply("do not edit!"); err != nil {
				log.Printf("failed to reply to message: %s", err)
			}
		case botgolang.CALLBACK_QUERY:
			data := update.Payload.CallbackQuery()
			switch data.CallbackData {
			case "echo":
				response := bot.NewButtonResponse(data.QueryID, "", "Hello World!", false)
				if err := response.Send(); err != nil {
					log.Printf("failed to reply on button click: %s", err)
				}
			}
		}
	}

}
