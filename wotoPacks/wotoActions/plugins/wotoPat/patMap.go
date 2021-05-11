// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoPat

import (
	"reflect"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// these two guys suppossed to be map, but I reconsider that,
// since it will took too much memory, I thought it would be
// better for them to be a pointer to a bson.D.
var patMap *[]bson.M
var patHMap *[]bson.M

// patIntit will initialize the pat map.
// it returns wa.SUCCESS if the map is intialized successfully,
// returns wa.FAILED, if something went wrong,
// and it returns wa.CANCELED, if map has already been initialized.
func patInit() wa.RESULT {
	if patMap == nil {
		settings := appSettings.GetExisting()
		if settings == nil {
			return wa.FAILED
		}

		client := settings.GetPatClient()
		if client == nil {
			return wa.FAILED
		}

		result, myMap := client.GetPatList()
		if result != wa.SUCCESS {
			return result
		}
		if myMap == nil {
			return wa.FAILED
		}

		patMap = &myMap
		return wa.SUCCESS
	}

	return wa.CANCELED
}

// patIntit will initialize the pat map.
// it returns wa.SUCCESS if the map is intialized successfully,
// returns wa.FAILED, if something went wrong,
// and it returns wa.CANCELED, if map has already been initialized.
func patHIntit() wa.RESULT {
	if patHMap == nil {
		settings := appSettings.GetExisting()
		if settings == nil {
			return wa.FAILED
		}

		client := settings.GetPatClient()
		if client == nil {
			return wa.FAILED
		}

		result, myMap := client.GetHPatList()
		if result != wa.SUCCESS {
			return result
		}
		if myMap == nil {
			return wa.FAILED
		}

		patHMap = &myMap
		return wa.SUCCESS
	}

	return wa.CANCELED
}

func patExists(ID string) bool {
	if patMap == nil {
		return false
	}

	var id string

	for _, current := range *patMap {
		if reflect.TypeOf(id) != reflect.TypeOf(current) {
			continue
		}
		id = current[wotoValues.PAT_ID_KEY].(string)
		if id == ID {
			return true
		}
	}

	return false
}

func patArrayExists(photo []tgbotapi.PhotoSize) bool {
	for _, current := range photo {
		if patExists(current.FileID) {
			return true
		}
	}
	return false
}

func patHExists(ID string) bool {
	if patHMap == nil {
		return false
	}

	var id string

	for _, current := range *patHMap {
		if reflect.TypeOf(id) != reflect.TypeOf(current) {
			continue
		}
		id = current[wotoValues.PAT_ID_KEY].(string)
		if id == ID {
			return true
		}
	}

	return false
}

func patArrayHExists(photo []tgbotapi.PhotoSize) bool {
	for _, current := range photo {
		if patHExists(current.FileID) {
			return true
		}
	}
	return false
}

func setNewPat(p *[]bson.M) {
	patMap = p
}

func setNewHPat(hp *[]bson.M) {
	patHMap = hp
}
