package main

import (
	"gopkg.in/ini.v1"
	telegramBot "gopkg.in/telebot.v3"
	"log"
	"tel-mtproto/common"
	"time"
)

var isIniInitOnce = false
var IniData *ini.File

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

	bot.Handle("/add_receiver", func(context telegramBot.Context) error {
		msg := context.Send("Send me the userIds you want to add to the list.\n" +
			"The format Should be as following:\n" +
			"Id1:name1\n" +
			"Id2:name2\n\n" +
			"It's better for name to not contain '-'.")
		println(msg)
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
