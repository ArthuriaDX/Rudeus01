package interfaces

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"
)

type WClient interface {
	AccountsLength() int
	PingClientDB(_report bool)
	DeleteCollection(database dbTypes.DATABASE, collection dbTypes.COLLECTION) common.RESULT
	Destroy()
	GetWotoConfiguration() common.RESULT
	ResetWotoConfiguration() common.RESULT
	ResetUsersCollection() common.RESULT
	CreateNewConfiguration() common.RESULT
	UpdateLastUIDConfiguration(_index uint8) common.RESULT
	CreateNewAccount() common.RESULT
	FindAccount(_username, _pass *string) common.RESULT
	DeleteAccount() common.RESULT
	UpdateAccount() common.RESULT
	UpdateAccountOnlineToken() common.RESULT
}
