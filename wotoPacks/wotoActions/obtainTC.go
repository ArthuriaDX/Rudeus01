package wotoActions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func WObtainTC(_settings *appSettings.AppSettings) bool {

	values := map[string]string{"request_value": "<get_token>"}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	indexStr := os.Getenv(wotoValues.INDEX_KEY)
	if wotoSecurity.IsEmpty(&indexStr) {
		return true
	}
	index, _err := strconv.Atoi(indexStr)

	if _err != nil {
		log.Fatal(_err)
	}

	resp, err := http.Post("https://frosty-surf-5435.saogame.workers.dev/", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Fatal(err)
	}
	str := string(body)
	myInt, _err := strconv.Atoi(str)
	if _err != nil {
		return false
	}
	return myInt == index
}
