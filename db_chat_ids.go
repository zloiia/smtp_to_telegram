package main

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ChatIdsBot struct {
	bot *tgbotapi.BotAPI
	db *gorm.DB
}

type ChatIdsBot_ptr *ChatIdsBot


func CreateChatIds() ChatIdsBot_ptr {
	return &ChatIdsBot{
		bot: nil,
		db : nil,
	}
}

func (b *ChatIdsBot) Connect(telegram_bot_token string, db_dsn string) error {
	if (b == nil) {
		return errors.New("empty pointer")
	}
	bot, err := tgbotapi.NewBotAPI(telegram_bot_token)
	if err != nil {
		return err
	}
	b.bot = bot
	b.bot.Debug = false

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: db_dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	  }), &gorm.Config{})
	if err != nil {
		return err
	}

	b.db = db

	return nil
}

func (b *ChatIdsBot) Start() error {
	if (b == nil) {
		return errors.New("empty pointer")
	}
	if (b.bot == nil) {
		return errors.New("connect first")
	}
	update_config  := tgbotapi.NewUpdate(0)
	update_config.Timeout = 30
	update_config.Limit = 1

	updates := b.bot.GetUpdatesChan(update_config)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if !update.Message.IsCommand() {
			msg.Text = "Please send /help"
			b.bot.Send(msg)
			continue
		}

		switch update.Message.Command() {
		case "receive_email":

		}
	}


	return nil
}
