package tgConst

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	USER_NAME          = "$WOTO_USER_NAME"
	FIRST_NAME         = "$WOTO_FIRST_NAME"
	USER_NAME_MENTION  = "$WOTO_USERNAME_MENTION"
	FIRST_NAME_MENTION = "$WOTO_FIRSTNAME_MENTION"
)

func ReplaceDyn(value *string, msg *tg.Message) {
	if wotoStrings.IsEmpty(value) || msg == nil {
		return
	}
	*value = wotoMD.GetNormal(*value).ToString()
	var tmp, md, userMd string

	md = wotoMD.GetNormal(USER_NAME).ToString()
	if strings.Contains(*value, md) {
		userMd = wotoMD.GetNormal(msg.From.UserName).ToString()
		tmp = strings.ReplaceAll(*value, md, userMd)
		*value = tmp
	}

	md = wotoMD.GetNormal(FIRST_NAME).ToString()
	if strings.Contains(*value, md) {
		userMd = wotoMD.GetNormal(msg.From.UserName).ToString()
		tmp = strings.ReplaceAll(*value, md, userMd)
		*value = tmp
	}

	md = wotoMD.GetNormal(USER_NAME_MENTION).ToString()
	if strings.Contains(*value, md) {
		// msg.From.ID doesn't need to be repaired at all,
		// and also msg.From.UserName will be repaired inside of the
		// function, so there is no need to repair for it.
		tmp = wotoMD.GetUserMention(msg.From.UserName, msg.From.ID).ToString()
		tmp = strings.ReplaceAll(*value, md, tmp)
		*value = tmp
	}

	md = wotoMD.GetNormal(FIRST_NAME_MENTION).ToString()
	if strings.Contains(*value, md) {
		// msg.From.ID doesn't need to be repaired at all,
		// and also msg.From.FirstName will be repaired inside of the
		// function, so there is no need to repair for it.
		tmp = wotoMD.GetUserMention(msg.From.FirstName, msg.From.ID).ToString()
		tmp = strings.ReplaceAll(*value, md, tmp)
		*value = tmp
	}
	appSettings.GetExisting().SendSudo(*value)
}
