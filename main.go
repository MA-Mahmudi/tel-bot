package main

import (
	"gopkg.in/ini.v1"
	telegramBot "gopkg.in/telebot.v3"
	"log"
	"tel-mtproto/admin-bot"
	"tel-mtproto/common"
	"time"
)

var isIniInitOnce = false
var IniData *ini.File

/////////// nodemon --exec go run main.go --signal SIGTERM

func main() {
	pref := telegramBot.Settings{
		Token:  IniGet("", "TOKEN"),
		Poller: &telegramBot.LongPoller{Timeout: 60 * time.Second},
	}

	bot, err := telegramBot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(context telegramBot.Context) error {
		welcomeMsg := context.Send("This is your admin bot. Enjoy using it!")
		return welcomeMsg
	})

	admin_bot.AddReceiver(bot)

	bot.Handle("/remove_receiver", func(context telegramBot.Context) error {
		msg := context.Send("Enter the userId(s) you want to remove from the list.\n" +
			"The format should be as following:\n\n" +
			"Id1-Id2-Id3")

		return msg
	})

	bot.Handle("/custom_text", func(context telegramBot.Context) error {
		msg := context.Send("Send the text you want to send to the receivers.")

		return msg
	})

	bot.Start()
}

func IniSetup() {
	if !isIniInitOnce {
		var err error
		IniData, err = ini.Load("config.ini")
		common.IsErr(err, "Error loading .ini file")
		isIniInitOnce = true
	} else {
		println("initialized inis once")
	}
}

func IniGet(section string, key string) string {
	if IniData == nil {
		IniSetup()
	}
	return IniData.Section(section).Key(key).String()
}
