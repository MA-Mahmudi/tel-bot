package adminBot

import (
	telegramBot "gopkg.in/telebot.v3"
)

func AddReceiver(bot *telegramBot.Bot) {

	//ReceiversList := [...]struct {
	//	Id   string
	//	Name string
	//}{{}}

	bot.Handle("/add_receiver", func(context telegramBot.Context) error {
		msg := context.Send(
			"Enter the userId(s) you want to add to the list.\n" +
				"The format must be as following:\n\n" +
				"Add:\n" +
				"Id:Name, Id:Name, ...")

		bot.Handle(telegramBot.OnText, func(context telegramBot.Context) error {
			res := context.Text()
			if len(res) > 5 && (res[:3] == "add" || res[:3] == "Add" || res[:3] == "ADD") {
				//for i := 5; i <= len(res); i++ {
				//	if res[i] == ':' {
				//		res
				//	}
				//}
				msg := context.Send("Receivers' list updated...")
				return msg
			}
			return nil
		})

		return msg
	})
}
