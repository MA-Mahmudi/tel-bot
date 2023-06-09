package main

import (
	"gopkg.in/ini.v1"
	telegramBot "gopkg.in/telebot.v3"
	"log"
	adminBot "tel-mtproto/adminBot"
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

	adminBot.AddReceiver(bot)

	adminBot.RemoveReceiver(bot)

	adminBot.CustomText(bot)

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
