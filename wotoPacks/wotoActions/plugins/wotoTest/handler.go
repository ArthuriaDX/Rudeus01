package wotoTest

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TestCommandHandle is not completed yet!
func TestCommandHandle(message *tg.Message, args pTools.Arg) {
	settings := appSettings.GetExisting()
	//log.Println("In EVENT!")
	if settings == nil {
		return
	}
	var str string
	if len(args) <= wotoValues.BaseOneIndex {
		str = "no test session provided!"
	} else {
		str = "couldn't find test session : " + args[wotoValues.BaseOneIndex]
	}
	msg := tg.NewMessage(message.Chat.ID, str)
	msg.ReplyToMessageID = message.MessageID
	if _, err := settings.GetAPI().Send(msg); err != nil {
		log.Println(err)
		return
	}
}
