// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoDB

import (
	"log"
	"reflect"

	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"go.mongodb.org/mongo-driver/bson"
)

// GetHPatList will give you the pat list in the db.
// Even if there is no pat in the db, `[]bson.M` will not be nil,
// though its length will be zero.
func (_c *WotoClient) GetPatList() (wa.RESULT, []bson.M) {
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}

	re, raws := _c.getRAWS(PatDataBase, PatListCollection)
	if re != wa.SUCCESS || raws == nil {
		log.Println(wv.GET_PAT_LIST_ERROR)
		return wa.FAILED, nil
	}
	return wa.SUCCESS, raws
}

// GetHPatList will give you the hentai pat list in the db.
// Even if there is no pat in the db, `[]bson.M` will not be nil,
// though its length will be zero.
func (_c *WotoClient) GetHPatList() (wa.RESULT, []bson.M) {
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}

	re, raws := _c.getRAWS(PatDataBase, HPatListCollection)
	if re != wa.SUCCESS || raws == nil {
		log.Println(wv.GET_PAT_LIST_ERROR)
		return wa.FAILED, nil
	}
	return wa.SUCCESS, raws
}

// AddPat will add a new pat to the db.
// it's not our duty to check if the patID already exists or not!
// pat plugin should check if before call this method!
func (_c *WotoClient) AddPat(patID string, t int32) (wa.RESULT, []bson.M) {
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}

	collection := _c.getCollection(PatDataBase, PatListCollection)

	b := bson.D{
		{Key: wv.PAT_ID_KEY, Value: patID},
		{Key: wv.PAT_TYPE_KEY, Value: t},
	}
	_, err := collection.InsertOne(*_c.ctx, b)

	if err != nil {
		log.Println(err)
		return wa.FAILED, nil
	}

	return _c.GetPatList()
}

// AddPat will add a new hentai pat to the db.
// it's not our duty to check if the patID already exists or not!
// pat plugin should check if before call this method!
func (_c *WotoClient) AddHPat(patID string, t int32) (wa.RESULT, []bson.M) {
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}

	collection := _c.getCollection(PatDataBase, HPatListCollection)

	b := bson.D{
		{Key: wv.PAT_ID_KEY, Value: patID},
		{Key: wv.PAT_TYPE_KEY, Value: t},
	}
	_, err := collection.InsertOne(*_c.ctx, b)

	if err != nil {
		log.Println(err)
		return wa.FAILED, nil
	}

	return _c.GetHPatList()
}

// RemovePat will remove the pat with specified id.
// It will return `FAILED` if it encouners any error,
// will return `SUCCESS` if it successfully deleted the pat and
// the call to getPatList was successful,
// and will return `CANCELED` if it found nothing in loop.
// It's not that it will give you information about existing of a pat!
// so be carefull about it; the pat MAY EXISTS in the db, but
// the method may return `CANCELED`.
func (_c *WotoClient) RemovePat(patID string) (wa.RESULT, []bson.M) {
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}

	collection := _c.getCollection(PatDataBase, PatListCollection)
	re, raws := _c.getRAWS(PatDataBase, PatListCollection)
	if re != wa.SUCCESS {
		log.Println("Couldn't get raw data from database in RemovePat method")
		return wa.FAILED, nil
	}

	id := wv.EMPTY
	for _, current := range raws {
		if reflect.TypeOf(current[wv.PAT_ID_KEY]) != reflect.TypeOf(id) {
			continue
		}
		id = current[wv.PAT_ID_KEY].(string)
		if id == patID {
			collection.DeleteOne(*_c.ctx, current)
			return _c.GetPatList()
		}
	}
	return wa.CANCELED, nil
}

// RemoveHPat will remove the hentai pat with specified id.
// It will return `FAILED` if it encouners any error,
// will return `SUCCESS` if it successfully deleted the pat and
// the call to getPatList was successful,
// and will return `CANCELED` if it found nothing in loop.
// It's not that it will give you information about existing of a pat!
// so be carefull about it; the pat MAY EXISTS in the db, but
// the method may return `CANCELED`.
func (_c *WotoClient) RemoveHPat(patID string) (wa.RESULT, []bson.M) {
	if !_c.isPatClient() {
		return wa.FAILED, nil
	}

	collection := _c.getCollection(PatDataBase, HPatListCollection)
	re, raws := _c.getRAWS(PatDataBase, HPatListCollection)
	if re != wa.SUCCESS {
		log.Println("Couldn't get raw data from database in RemovePat method")
		return wa.FAILED, nil
	}

	id := wv.EMPTY
	for _, current := range raws {
		if reflect.TypeOf(current[wv.PAT_ID_KEY]) != reflect.TypeOf(id) {
			continue
		}
		id = current[wv.PAT_ID_KEY].(string)
		if id == patID {
			collection.DeleteOne(*_c.ctx, current)
			return _c.GetPatList()
		}
	}
	return wa.CANCELED, nil
}
