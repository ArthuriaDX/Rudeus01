// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoQuery

import (
	"encoding/json"
	"errors"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	NoneQueryPlugin QueryPlugin = 0 // None
	UdQueryPlugin   QueryPlugin = 1 // wotoUD plugin
)

type QueryPlugin uint8

type QueryBase struct {
	Plugin   int    `json:"p"` // the plugin
	Data     int    `json:"d"` // the data (unique data code)
	UniqueId string `json:"u"` // the unique id of source message
}

type WotoQuery interface {
	GetId() int
	GetUniqueId() string
}

var queryMap map[int]WotoQuery
var last int = 0

func ToQueryBase(data string) (base *QueryBase, err error) {
	if wotoStrings.IsEmpty(&data) {
		return nil, errors.New(dataEmptyErr)
	}

	b := QueryBase{}
	errJson := json.Unmarshal([]byte(data), &b)
	if errJson != nil {
		return nil, errJson
	}

	return &b, nil
}

func GetQuery(plugin QueryPlugin, dataCode int, unique string) *QueryBase {
	return &QueryBase{
		Plugin:   int(plugin),
		Data:     dataCode,
		UniqueId: unique,
	}
}

func SetInMap(value WotoQuery) int {
	mapInit()
	last++
	for {
		if last == wv.BaseIndex {
			return wv.BaseIndex
		}
		if queryMap[last] != nil {
			last++
			continue
		}
		queryMap[last] = value
		return last
	}
}

// RemoveFromMap will remove all of the queries with the specified
// id from the map.
// the ID should be message id.
func RemoveFromMap(ID int) {
	for i := wv.BaseIndex; i <= last; i++ {
		if queryMap[i] != nil {
			if queryMap[i].GetId() == ID {
				queryMap[i] = nil
			}
		}
	}
}

func mapInit() {
	if queryMap == nil {
		queryMap = make(map[int]WotoQuery)
	}
}

func getFromMap(index int, uniqueId string) WotoQuery {
	current := queryMap[index]
	if current == nil {
		return current
	}

	if current.GetUniqueId() != uniqueId {
		return nil
	}

	return current
}
