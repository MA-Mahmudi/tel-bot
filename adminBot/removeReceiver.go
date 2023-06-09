package adminBot

import telegramBot "gopkg.in/telebot.v3"

func RemoveReceiver(bot *telegramBot.Bot) {

	bot.Handle("/remove_receiver", func(context telegramBot.Context) error {
		msg := context.Send("Enter the userId(s) you want to remove from the list.\n" +
			"The format should be as following:\n\n" +
			"Id1-Id2-Id3")

		return msg
	})

}
