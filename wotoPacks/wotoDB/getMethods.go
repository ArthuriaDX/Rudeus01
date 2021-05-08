// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import (
	"log"

	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB/dbTypes"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

func (_c *WotoClient) GetWotoConfiguration() wa.RESULT {
	re, raws := _c.getRAWS(MainDataBase, ConfigurationCollection)
	if re != wa.SUCCESS {
		log.Println("Couldn't get raw in GetWotoConfiguration method!")
		return wa.FAILED
	}

	mainSudo := raws[wotoValues.BaseIndex][MAIN_SUDO_KEY].(int64)
	if mainSudo == wotoValues.BaseIndex {
		log.Println("Tried to get main sudo id from db, but it is zero!")
	}

	sudoL := raws[wotoValues.BaseIndex][SUDO_LIST_KEY].(string)
	if ws.IsEmpty(&sudoL) {
		log.Println("Tried to get sudo list from db, but it's Empty!")
	}

	_c.settings.SetMainSudo(mainSudo)
	_c.settings.SetSudoList(sudoL)

	//appSettings.GetExisting().SendSudo(strconv.FormatInt(mainSudo, wotoValues.BaseTen))

	//var currentValue string
	//for _, current := range raws {
	//	currentValue := current["AdminList"].(string)
	//	if
	//}
	//str := raws[0]["AdminList"].(string)
	//appSettings.GetExisting().SendSudo(str)
	//_w := wotoConfiguration{
	//	UIDKeyName: [MaxUIDIndex]string{},
	//	LastUID:    [MaxUIDIndex]UID{},
	//}
	//_re, _raws := _c.getRAWS(MainDataBase, ConfigurationCollection)
	//if _re != SUCCESS {
	//	return FAILED, nil
	//}
	//_w.setUIDKeys()
	//for _i, _current := range _raws {
	//	_value, _ok := _current[UIDKeyName].(string)
	//	_w.LastUID[_i] = UID(_value)
	//	if !_ok {
	//		clientError()
	//		return FAILED, nil
	//	}
	//}
	return wa.SUCCESS
}

// getRAWS will give you the raw data.
func (_c *WotoClient) getRAWS(database dbTypes.DATABASE, collection dbTypes.COLLECTION) (wa.RESULT, []bson.M) {
	_collection := _c.getCollection(database, collection)
	if _collection == nil {
		// clientError()
		// well, it seems it's really rare to reach this poit,
		// but I will note this so do something about it in the future.
		return wa.FAILED, nil
	}
	_cursor, _cursorError := _collection.Find(*_c.ctx, bson.M{})
	if _cursorError != nil || _cursor == nil {
		return wa.FAILED, nil
	}
	var _raws []bson.M
	_ = _cursor.All(*_c.ctx, &_raws)
	return wa.SUCCESS, _raws
}
