package main

import (
    i "brial-bot-go/internal"
)

func main() {
    bot, err := i.NewBot()

    if err != nil {
        panic(err)
    }

    bot.Run()
}
