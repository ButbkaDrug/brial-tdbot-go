package internal

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Bot struct {
    api *tgbotapi.BotAPI
    bri *bri
}

func NewBot() (*Bot, error){

    if err := godotenv.Load(".env"); err != nil {
        panic("Cant load enviroment" + err.Error())
    }

    key := os.Getenv("BRIAL_BOT_API_KEY")

    bot, err := tgbotapi.NewBotAPI(key)

    if err != nil {
        return nil, err
    }

    return &Bot{
        api: bot,
        bri: NewBri(),
    }, nil
}

func(b *Bot) Run() {
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := b.api.GetUpdatesChan(u)

    for update := range updates {

        b.updateHandler(update)
    }
}

func(b *Bot) updateHandler(u tgbotapi.Update) {
    var response string

    if u.Message == nil || u.Message.Text == "" {return}



    msg := u.Message.Text

    if b.bri.IsEncoded(msg) {
        response = b.bri.Decode(msg)
    } else {
        response = b.bri.Encode(msg)
    }

    message := tgbotapi.NewMessage(u.Message.Chat.ID, response)
    _, err := b.api.Send(message)

    if err != nil {
        log.Println(err)
    }
}
