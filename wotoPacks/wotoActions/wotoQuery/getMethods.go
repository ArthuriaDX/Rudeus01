package wotoQuery

import (
	"encoding/json"
	"log"

	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func (q *QueryBase) Exists() bool {
	mapInit()

	m := queryMap[q.Data]

	return m != nil
}

func (q *QueryBase) ToString() string {
	b, err := json.Marshal(*q)
	if err != nil {
		log.Println(err)
		return wv.EMPTY
	}

	if b == nil || len(b) == wv.BaseIndex {
		return wv.EMPTY
	}

	return string(b)
}

func (q *QueryBase) GetType() QueryPlugin {
	return QueryPlugin(q.Plugin)
}

func (q *QueryBase) GetDataQuery() WotoQuery {
	return getFromMap(q.Data)
}
