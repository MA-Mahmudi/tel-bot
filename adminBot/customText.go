package adminBot

import telegramBot "gopkg.in/telebot.v3"

func CustomText(bot *telegramBot.Bot) {

	bot.Handle("/custom_text", func(context telegramBot.Context) error {
		msg := context.Send("Send the text you want to send to the receivers.")

		return msg
	})

}
