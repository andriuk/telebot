package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

// Telegram bot
var (
	TeleToken = os.Getenv("TELEGRAM_TOKEN")
)

// kbotCmd represents the kbot command

var kbotCmd = &Command{
	Use:   "kbot",
	Short: "kbot is a telegram bot",
	Long:  `kbot is a telegram bot`,

	Run: func(cmd *cobra.command, args []string) {

		fmt.Printf("kbot %s started", appVersion)
		log.Println("kbot is running")

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatal(err)
			log.Fatalf("Please check TELEGRAM_TOKEN env variable %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			return err
		})

		kbot.Start()

	},
}
