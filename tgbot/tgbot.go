package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func tgbot() {
	bot, err := tgbotapi.NewBotAPI("6203088041:AAFgPoZgWlZvkAFP1KWK1kd2DtOXV3AXHDM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "Йоу Братишка я меняю размер картиночки вызывай команду /redactimage и я ебану всё в лучшем виде!"
		case "redactimage":
			msg.Text = "Кинь ссылку на изображение"
			//file := tgbotapi.FileURL("https://i.imgur.com/unQLJIb.jpg")

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
