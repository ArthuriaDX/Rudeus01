// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import (
	"log"
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

func (_c *WotoClient) CreateNewConfiguration() wa.RESULT {
	collection := _c.getCollection(MainDataBase, ConfigurationCollection)
	SudoList := strconv.FormatInt(wv.SUDO, wv.BaseTen)
	b := bson.D{
		{Key: MAIN_SUDO_KEY, Value: wv.SUDO},
		{Key: SUDO_LIST_KEY, Value: SudoList},
	}

	_, err := collection.InsertOne(*_c.ctx, b)
	if err != nil {
		appSettings.GetExisting().SendSudo(err.Error())
	}

	return wa.SUCCESS
}

func (_c *WotoClient) AddSudo(id int64) wa.RESULT {

	if id == wv.BaseIndex {
		log.Println("The ID cannot be zero!")
		return wa.FAILED
	}

	if _c.settings.IsSudo(id) {
		log.Println("This user is already a sudo!")
		return wa.CANCELED
	}

	collection := _c.getCollection(MainDataBase, ConfigurationCollection)

	re, raws := _c.getRAWS(MainDataBase, ConfigurationCollection)
	if re != wa.SUCCESS {
		log.Println("Couldn't get raw data from database in AddSudo method")
		return wa.FAILED
	}

	sudoList := raws[wv.BaseIndex][SUDO_LIST_KEY].(string)
	sudoList += wv.SEMICOLON + strconv.FormatInt(id, wv.BaseTen)
	b := bson.D{
		{Key: wv.SET_VALUE, Value: bson.D{
			{Key: MAIN_SUDO_KEY, Value: wv.SUDO},
			{Key: SUDO_LIST_KEY, Value: sudoList},
		}},
	}

	_, err := collection.UpdateOne(*_c.ctx, raws[wv.BaseIndex], b)
	if err != nil {
		log.Println(err)
		appSettings.GetExisting().SendSudo(err.Error())
		return wa.FAILED
	}

	_c.GetWotoConfiguration()

	return wa.SUCCESS
}

func (_c *WotoClient) RemSudo(id int64) wa.RESULT {

	if id == wv.BaseIndex {
		log.Println("ID cannot be zero!")
		return wa.FAILED
	}

	if !_c.settings.IsSudo(id) || _c.settings.IsMainSudo(id) {
		log.Println("Connot remove main sudo!")
		return wa.CANCELED
	}

	collection := _c.getCollection(MainDataBase, ConfigurationCollection)

	re, raws := _c.getRAWS(MainDataBase, ConfigurationCollection)
	if re != wa.SUCCESS {
		log.Println("Tried to get raw data from db, but couldn't!")
		return wa.FAILED
	}

	l := _c.settings.GetSudoList()
	sudoList := strconv.FormatInt(wv.SUDO, wv.BaseTen)

	for _, current := range l {
		if current == id {
			continue
		}

		if current == wv.SUDO {
			continue
		}

		sudoList += wv.SEMICOLON + strconv.FormatInt(current, wv.BaseTen)
	}
	//sudoList := raws[wv.BaseIndex][SUDO_LIST_KEY].(string)
	// sudoList := raws[wv.BaseIndex][MAIN_SUDO_KEY].(string)
	b := bson.D{
		{Key: wv.SET_VALUE, Value: bson.D{
			{Key: MAIN_SUDO_KEY, Value: wv.SUDO},
			{Key: SUDO_LIST_KEY, Value: sudoList},
		}},
	}

	_, err := collection.UpdateOne(*_c.ctx, raws[wv.BaseIndex], b)
	if err != nil {
		log.Println(err)
		appSettings.GetExisting().SendSudo(err.Error())
		return wa.FAILED
	}

	_c.GetWotoConfiguration()

	return wa.SUCCESS
}
