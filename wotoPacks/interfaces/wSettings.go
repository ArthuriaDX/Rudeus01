package interfaces

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Go encourages composition over inheritance, using simple, often one-method interfaces ...
// that serve as clean, comprehensible boundaries between components.
// —Rob Pike,
//	“Go at Google: Language Design in the
//	Service of Software Engineering”
//	(see talks.golang.org/ 2012/splash.article)

type WSettings interface {
	SetAPI(_api *tg.BotAPI)
	SetObt(_obt func(*WSettings) bool)
	SetNextChild(_nextRunner func(string, *WSettings))
	SetURL(_url string)
	SetNextURL(_url string)
	SetTObt(_obt string)
	SetIndex(_index int)
	InvalidateAPI() bool
	App_enter()
	App_exit()
	RunNext()
	SendSudo(str string)
}
