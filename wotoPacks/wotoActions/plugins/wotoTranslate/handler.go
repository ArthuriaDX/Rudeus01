package wotoTranslate

import (
	"log"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TrHandler(message *tg.Message, args pTools.Arg) {

}

func ToMorseHandler(message *tg.Message, args pTools.Arg) {
	// set the first element of args to empty,
	// because pTools.Arg also contains the command itself,
	// however we don't want that here.
	args[wv.BaseIndex] = wv.EMPTY

	is_bin := args.HasFlag(BIN_FLAG, BINARY_FLAG)
	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)

	is_reply := message.ReplyToMessage != nil
	var full, trl string
	if is_reply {
		if !ws.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			// we should check if the replied message has text or not,
			// and since it doesn't have any text value, then we should not
			// do any operations on it.
			return
		}
	} else {
		// do not convert the flags to the morse code.
		full = args.JoinNoneFlags(false)
	}

	if is_bin {
		trl = TranslateToBinary(full)
	} else {
		trl = TranslateToMorse(full)
	}
	var str string
	if !ws.IsEmpty(&trl) {
		if is_bin {
			md1 := wotoMD.GetBold(TO_BINARY_M)
			md2 := wotoMD.GetNormal(trl)
			str = md1.Append(md2).ToString()
		} else {
			md1 := wotoMD.GetBold(TO_MORSE_M)
			md2 := wotoMD.GetNormal(trl)
			str = md1.Append(md2).ToString()
		}
	}
	sendMorse(message, &str, is_reply, send_pv)
}

func FromMorseHandler(message *tg.Message, args pTools.Arg) {
	// set the first element of args to empty,
	// because pTools.Arg also contains the command itself,
	// however we don't want that here.
	args[wv.BaseIndex] = wv.EMPTY

	is_reply := message.ReplyToMessage != nil
	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)

	var full string
	if is_reply {
		if !ws.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			return
		}
	} else {
		full = strings.Join(args, wv.SPACE_VALUE)
	}
	trl := TranslateFromMorse(full)
	var str string
	if !ws.IsEmpty(&trl) {
		md := wotoMD.GetBold(FROM_MORSE_M)
		normal := wotoMD.GetNormal(trl)
		str = md.Append(normal).ToString()
		//str = wotoMD.GetBold(FROM_MORSE_M).ToString()
		//str += wotoMD.GetNormal(trl).ToString()
	}
	sendMorse(message, &str, is_reply, send_pv)
}

// please before using this function, ensure that the `morse` value is
// converted to markdown (using wotoMarkdown) functions.
func sendMorse(message *tg.Message, morse *string, reply, pv bool) {
	// always check before sending:
	// Bad Request: message text is empty
	if ws.IsEmpty(morse) {
		return
	}

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	var msg tg.MessageConfig
	if pv {
		//settings.SendSudo(strconv.Itoa(int(message.From.ID)))
		msg = tg.NewMessage(message.From.ID, (*morse))
	} else {
		//settings.SendSudo(strconv.Itoa(int(message.Chat.ID)))
		msg = tg.NewMessage(message.Chat.ID, (*morse))
	}

	// !pv => for fixing: Bad Request: replied message not found
	if reply && !pv {
		r := message.ReplyToMessage
		if r != nil {
			msg.ReplyToMessageID = r.MessageID
		} else {
			msg.ReplyToMessageID = message.MessageID
		}
	} else {
		// for fixing: Bad Request: replied message not found
		if !pv {
			msg.ReplyToMessageID = message.MessageID
		}
	}
	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		tgErr.SendRandomErrorMessage(message)
		return
	}
}
