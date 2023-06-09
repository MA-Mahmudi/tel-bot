package adminBot

import telegramBot "gopkg.in/telebot.v3"

func AddReceiver(bot *telegramBot.Bot) {

	bot.Handle("/add_receiver", func(context telegramBot.Context) error {
		msg := context.Send(
			"Enter the userId(s) you want to add to the list.\n" +
				"The format should be as following:\n" +
				"Add:\n" +
				"Id1:name1, Id2:name2, ...\n\n" +
				"It's better for name to not contain '-'.")

		bot.Handle(telegramBot.OnText, func(context telegramBot.Context) error {
			res := context.Text()
			if len(res) > 5 && (res[:3] == "add" || res[:3] == "Add" || res[:3] == "ADD") {
				msg := context.Send("Receivers' list updated...")
				return msg
			}
			return nil
		})

		return msg
	})

	//bot.Handle(telegramBot.OnText, func(context telegramBot.Context) error {
	//	res := context.Text()
	//	if len(res) > 0 {
	//		msg := context.Send("Ok")
	//		return msg
	//	}
	//	return nil
	//})

}
