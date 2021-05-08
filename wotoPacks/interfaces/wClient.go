package interfaces

import (
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"
)

type WClient interface {
	AccountsLength() int
	PingClientDB(_report bool)
	DeleteCollection(database dbTypes.DATABASE, collection dbTypes.COLLECTION) wa.RESULT
	Destroy()
	GetWotoConfiguration() wa.RESULT
	ResetWotoConfiguration() wa.RESULT
	ResetUsersCollection() wa.RESULT
	CreateNewConfiguration() wa.RESULT
	CreateNewAccount() wa.RESULT
	AddSudo(id int64) wa.RESULT
	RemSudo(id int64) wa.RESULT
	FindAccount(_username, _pass *string) wa.RESULT
	DeleteAccount() wa.RESULT
	UpdateAccount() wa.RESULT
	UpdateAccountOnlineToken() wa.RESULT
}
