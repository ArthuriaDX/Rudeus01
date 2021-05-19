package wotoTranslate

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrong"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

//import "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoTranslate/wotoLang"

func DetectLanguage(text string) *Lang {
	m := map[string]string{
		userAgentKey:      userAgentValue,
		acceptKey:         acceptValue,
		acceptLanguageKey: acceptLanguageValue,
		refererKey:        refererValue,
		contentTypeKey:    contentTypeValue,
		originKey:         originValue,
		connectionKey:     connectionValue,
		teKey:             teValue,
		qKey:              text,
	}

	data, errJ := json.Marshal(m)
	if errJ != nil {
		log.Fatal(errJ)
	}

	reader := bytes.NewReader(data)
	resp, errH := http.Post(dHostUrl, contentTypeValue, reader)

	if errH != nil {
		log.Fatal(errH)
	}

	defer resp.Body.Close()

	b, errB := ioutil.ReadAll(resp.Body)
	if errB != nil {
		log.Fatal(errB)
	}

	str := ws.Qsb(b)
	strs := str.SplitStr(preLeft, preRight)
	log.Println(strs)

	b = []byte(strs[wv.BaseOneIndex].GetValue())

	var l Lang
	errJ = json.Unmarshal(b, &l)
	if errJ != nil {
		log.Fatal(errJ)
	}

	return &l
}
