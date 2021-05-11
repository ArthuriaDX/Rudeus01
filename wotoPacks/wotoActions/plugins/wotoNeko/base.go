// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	hostUrl     = "https://nekos.life/api/v2"
	headPatUrl  = "https://headp.at"
	slash       = "/"
	pats        = "/pats/"
	postContent = "application/json"
	patJs       = "js/pats.json"
	img         = hostUrl + "/img/"
	owo         = hostUrl + slash + OwoText + "?text="
	jpg         = ".jpg"
	jpeg        = ".jpeg"
	gif         = ".gif"
	png         = ".png"
)

// headp.at headers values.
// WARNING: do NOT edit these values if you
// don't know what the fuck is happening here.
const (
	userAgentKey = "User-Agent"
	fakeBrowser  = "Mozilla/5.0 "
	fakeOs       = "(X11; U; Linux i686) "
	gecko        = "Gecko/20071127 "
	firefox      = "Firefox/2.0.0.11"
)

type NekoBase struct {
	Url  string `json:"url"`
	Why  string `json:"why"`
	Name string `json:"name"`
	Cat  string `json:"cat"`
	Fact string `json:"fact"`
	Owo  string `json:"owo"`
}

// GetRandomSfw will give you a random sfw nekobase.
func GetRandomSfw() (base *NekoBase, err error) {
	ep := getRandomSfw()
	return getRandomThing(string(ep))
}

// GetRandomNsfw will give you a random nsfw nekobase.
func GetRandomNsfw() (base *NekoBase, err error) {
	ep := getRandomNSfw()
	return getRandomThing(string(ep))
}

// GetRandomTickle will give you a random tickle nekobase.
func GetRandomTickle() (base *NekoBase, err error) {
	return getRandomThing(string(Tickle))
}

// GetRandomSlap will give you a random slap nekobase.
func GetRandomSlap() (base *NekoBase, err error) {
	return getRandomThing(string(Slap))
}

// GetRandomPoke will give you a random poke nekobase.
func GetRandomPoke() (base *NekoBase, err error) {
	return getRandomThing(string(Poke))
}

// GetRandomPat will give you a random pat nekobase.
func GetRandomPat() (base *NekoBase, err error) {
	return getRandomThing(string(Pat))
}

// GetHeadPat will give you a random headpat, using headp.at api.
func GetHeadPat() (base *NekoBase, err error) {
	addr := headPatUrl + slash + patJs
	req, errJ := json.Marshal(map[string]string{
		userAgentKey: fakeBrowser + fakeOs + gecko + firefox,
	})
	if errJ != nil {
		log.Println(errJ)
		return nil, errJ
	}

	resp, err := http.Post(addr, postContent, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		// log.Println("In Resp.Body ERR!" + err.Error())
		return nil, errBody
	}

	array := make([]string, wv.BaseIndex)

	err = json.Unmarshal(b, &array)
	if err != nil {
		return nil, err
	}
	if len(array) == wv.BaseIndex {
		return nil, errors.New(headpatArrEmpty)
	}

	url := array[rand.Intn(len(array))]
	if ws.IsEmpty(&url) {
		return nil, errors.New(headpatUrlEmpty)
	}

	url = headPatUrl + pats + url
	neko := NekoBase{
		Url: url,
	}

	return &neko, nil
}

// GetRandomNeko will give you a random neko nekobase.
func GetRandomNeko() (base *NekoBase, err error) {
	return getRandomThing(string(Neko))
}

// GetRandomMeow will give you a random meow nekobase.
func GetRandomMeow() (base *NekoBase, err error) {
	return getRandomThing(string(Meow))
}

// GetRandomLizard will give you a random lizard nekobase.
func GetRandomLizard() (base *NekoBase, err error) {
	return getRandomThing(string(Lizard))
}

// GetRandomKiss will give you a random kiss nekobase.
func GetRandomKiss() (base *NekoBase, err error) {
	return getRandomThing(string(Kiss))
}

// GetRandomHug will give you a random hug nekobase.
func GetRandomHug() (base *NekoBase, err error) {
	return getRandomThing(string(Hug))
}

// GetRandomFoxGirl will give you a random fox girl nekobase.
func GetRandomFoxGirl() (base *NekoBase, err error) {
	return getRandomThing(string(Fox_Girl))
}

// GetRandomFeed will give you a random feed nekobase.
func GetRandomFeed() (base *NekoBase, err error) {
	return getRandomThing(string(Feed))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomCuddle() (base *NekoBase, err error) {
	return getRandomThing(string(Cuddle))
}

// GetRandomKemonomimi will give you a random cuddle nekobase.
func GetRandomKemonomimi() (base *NekoBase, err error) {
	return getRandomThing(string(Kemonomimi))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomHolo() (base *NekoBase, err error) {
	return getRandomThing(string(Holo))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomSmug() (base *NekoBase, err error) {
	return getRandomThing(string(Smug))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomBaka() (base *NekoBase, err error) {
	return getRandomThing(string(Baka))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomWoof() (base *NekoBase, err error) {
	return getRandomThing(string(Woof))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomGoose() (base *NekoBase, err error) {
	return getRandomThing(string(Goose))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomGecg() (base *NekoBase, err error) {
	return getRandomThing(string(Gecg))
}

// GetRandomCuddle will give you a random cuddle nekobase.
func GetRandomAvatar() (base *NekoBase, err error) {
	return getRandomThing(string(Avatar))
}

// GetRandomWaifu will give you a random waifu nekobase.
func GetRandomWaifu() (base *NekoBase, err error) {
	return getRandomThing(string(Waifu))
}

// GetRandomWhy will give you a random why nekobase.
func GetRandomWhy() (base *NekoBase, err error) {
	return getRandomText(WhyText)
}

// GetRandomName will give you a random name nekobase.
func GetRandomName() (base *NekoBase, err error) {
	return getRandomText(NameText)
}

// GetRandomCat will give you a random cat nekobase.
func GetRandomCat() (base *NekoBase, err error) {
	return getRandomText(CatText)
}

// GetRandomFact will give you a random fact nekobase.
func GetRandomFact() (base *NekoBase, err error) {
	return getRandomText(FactText)
}

// GetOwo will give you a random owo nekobase.
func GetOwo(value string) (base *NekoBase, err error) {
	if ws.IsEmpty(&value) {
		return nil, errors.New(owoEmpty)
	}

	addr := owo + value
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	b, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		// log.Println("In Resp.Body ERR!" + err.Error())
		return nil, errBody
	}

	neko := NekoBase{}
	err = json.Unmarshal(b, &neko)
	if err != nil {
		// log.Println("In UnMarshal ERR!" + err.Error())
		return nil, err
	}

	// log.Println("Final neko url is:" + neko.Url)
	return &neko, nil
}

// getRandomSfw will give you a random sfw endpoint.
func getRandomSfw() sfwNeko {
	initSfw()
	return sfw[rand.Intn(len(sfw))]
}

// getRandomNSfw will give you a random nsfw endpoint.
func getRandomNSfw() nsfwNeko {
	initNsfw()
	return nsfw[rand.Intn(len(nsfw))]
}

// getRandomThing will give you a random thing nekobase. LOL.
func getRandomThing(endpoint string) (base *NekoBase, err error) {
	addr := img + endpoint
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		// log.Println("In Resp.Body ERR!" + err.Error())
		return nil, errBody
	}

	neko := NekoBase{}
	err = json.Unmarshal(b, &neko)
	if err != nil {
		// log.Println("In UnMarshal ERR!" + err.Error())
		return nil, err
	}

	// log.Println("Final neko url is:" + neko.Url)
	return &neko, nil
}

// getRandomText will give you a random text nekobase. LOL.
func getRandomText(endpoint string) (base *NekoBase, err error) {
	addr := hostUrl + slash + endpoint
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		// log.Println("In Resp.Body ERR!" + err.Error())
		return nil, errBody
	}

	neko := NekoBase{}
	err = json.Unmarshal(b, &neko)
	if err != nil {
		// log.Println("In UnMarshal ERR!" + err.Error())
		return nil, err
	}

	// log.Println("Final neko url is:" + neko.Url)
	return &neko, nil
}
