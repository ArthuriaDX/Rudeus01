package wotoUD

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//---------------------------------------------------------

//---------------------------------------------------------
func (o *udOrigin) setCollection(c *UrbanCollection) {
	if o.collection != c {
		o.collection = c
	}
}

func (o *udOrigin) setButtons(b *tg.InlineKeyboardMarkup) {
	if o.keyboard != b {
		o.keyboard = b
	}
}

//---------------------------------------------------------
