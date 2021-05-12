package callBackQ

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoUD"
	wq "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoQuery"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func QHandler(query *tg.CallbackQuery) {
	base, qErr := wq.ToQueryBase(query.Data)
	if qErr != nil {
		log.Println(qErr)
		return
	}

	if base == nil || !base.Exists() {
		cancelMessage(query)
		return
	}

	switch base.GetType() {
	case wq.NoneQueryPlugin:
		cancelMessage(query)
		return
	case wq.UdQueryPlugin:
		wotoUD.QUdHanler(query, base)
	default:
		return
	}
}

func cancelMessage(query *tg.CallbackQuery) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	chat := query.Message.Chat.ID
	msg := query.Message.MessageID
	config := tg.NewEditMessageText(chat, msg, query.Message.Text)
	config.Entities = query.Message.Entities
	_, err := api.Request(config)
	if err != nil {
		log.Println(err)
		return
	}
}
