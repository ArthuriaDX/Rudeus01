// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

import (
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"
	"go.mongodb.org/mongo-driver/bson"
)

type WClient interface {
	AccountsLength() int
	PingClientDB(_report bool)
	DeleteCollection(database dbTypes.DATABASE, collection dbTypes.COLLECTION) wa.RESULT
	Destroy()
	GetWotoConfiguration() wa.RESULT
	GetPatList() (wa.RESULT, []bson.M)
	GetHPatList() (wa.RESULT, []bson.M)
	ResetWotoConfiguration() wa.RESULT
	ResetUsersCollection() wa.RESULT
	CreateNewConfiguration() wa.RESULT
	CreateNewAccount() wa.RESULT
	AddSudo(id int64) wa.RESULT
	RemSudo(id int64) wa.RESULT
	AddPat(patID string, t int32) (wa.RESULT, []bson.M)
	AddHPat(patID string, t int32) (wa.RESULT, []bson.M)
	RemovePat(patID string) (wa.RESULT, []bson.M)
	RemoveHPat(patID string) (wa.RESULT, []bson.M)
	FindAccount(_username, _pass *string) wa.RESULT
	DeleteAccount() wa.RESULT
	UpdateAccount() wa.RESULT
	UpdateAccountOnlineToken() wa.RESULT
}
