// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoUD

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	unsignedStr1 = "\\u003cb\\u003e"
	unsignedStr2 = "\\u003c\\/b\\u003e"
	// unsignedStrs3 = "\u003c" don't use them separately
	// unsignedStrs4 = "\u003e" don't use them separately
)

func GetSimilarWords(word string) (str []string, err error) {
	if ws.IsEmpty(&word) {
		return nil, errors.New(emptyWordErr)
	}

	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = url.PathEscape(word)

	urlStr := getURL(word)
	if ws.IsEmpty(&urlStr) {
		return nil, errors.New(emptyUrlErr)
	}

	//log.Println(urlStr)
	resp, httpErr := http.Get(urlStr)
	if httpErr != nil {
		return nil, httpErr
	}

	defer resp.Body.Close()

	b, bError := ioutil.ReadAll(resp.Body)
	if bError != nil {
		return nil, bError
	}

	// log.Println(string(b))
	gStr := parseGoogleData(b)
	log.Println(gStr)

	return gStr, nil
}

// getURL will give you a correct url to send the http request.
// NOTICE: we won't fix the word in this function, you have to
// do it before calling this function. please use url.PathEscape(string)
// function before calling this function.
func getURL(word string) string {
	if strings.Contains(word, wv.DY_WOTO_TEXT) {
		word = strings.ReplaceAll(word, wv.DY_WOTO_TEXT, wv.WOTO_TEXT)
	}

	return strings.ReplaceAll(googleSimUrl, wv.DY_WOTO_TEXT, word)
}

// parseGoogleData will fix the google data.
func parseGoogleData(gData []byte) []string {
	gStr := string(gData)
	final := wv.EMPTY
	allStrs := make([]string, wv.BaseIndex)
	firstChild := false
	secondChild := false
	thirdChild := false
	insideStr := false
	childDone := false
	additionalChild := wv.BaseIndex

	// let me explain what should we do here.
	// if you look at example09.json files,
	// you will see something like:
	// ["\u003cb\u003ered door yellow door\u003c\/b\u003e", 0, [7, 10, 30]],
	// we don't want these unicode characters.
	// we have to remove them.
	// format of these characters is:
	// \uxxxx (4 bits)
	// here is a list of this characte
	// [more information=>
	// 		https://en.wikipedia.org/wiki/List_of_Unicode_characters ]
	// 	"\u003cb\u003e" => <b> (bold)
	// 	"\u003c\/b\u003e" => </b> (bold end)
	// 	'\u003c' => '<'
	// 	'\u003e' => '>'
	// you can see the purified texts in example09-pure.json
	gStr = strings.ReplaceAll(gStr, unsignedStr1, wv.EMPTY)
	gStr = strings.ReplaceAll(gStr, unsignedStr2, wv.EMPTY)

	for _, current := range gStr {
		// CHAR_S7  = '['
		if current != wotoMD.CHAR_S7 {
			// CHAR_STR = "\""
			if current == wv.CHAR_STR {
				// ensure that we are in correct time line, Steins;Gate.
				if !thirdChild {
					// Forbidden!
					// we are only allowed to do operations in
					// third child, not in any other childs!
					continue
				}

				if !childDone {
					if insideStr {
						insideStr = false
						childDone = true
					} else {
						insideStr = true
					}

					continue
				} else {
					continue
				}
			}

			// CHAR_S8  = ']'
			if current == wotoMD.CHAR_S8 {
				// ensure that we have already passed the first child.
				// actually I don't know what the hell are these characters
				// at the first of recieved data from google: )]}'.
				// really, what's wrong with them? maybe it's only me
				// whose response is like this.
				// anyway, it won't hurt if I write a checker here.
				if !firstChild {
					continue
				}

				// we should check if this ended child was
				// an official child or not.
				if additionalChild == wv.BaseIndex {
					// it means this is an official child.

					// we should check if this ended child
					// is the third child or not.
					if thirdChild {
						// it means third child has been ended.
						// now our duty is to append it to the
						// allStrs array.
						allStrs = append(allStrs, final)
						final = wv.EMPTY
						childDone = false
						thirdChild = false
						continue
					}

					if secondChild {
						// if the result is not found in google's server,
						// we will get a response like this:
						// [[],{"i":"rellowrellow","q":"kikrhTfn3q01Ssg4sD_3SYeA8B4"}]
						if len(allStrs) == wv.BaseIndex {

							return nil
						}

						// we are done! return all of the slices.
						return allStrs
					}

					// no need for check for first child,
					// since we don't know if its contents is
					// bullshit or not.
					// for more information, compare
					// example09.json and example08.json with each other.
					// if firstChild {

					//}
					continue
				} else {
					// this means that the unofficial child's counter,
					// is not reached to zero yet.
					// so it means this ended child, was an unofficial
					// child.
					additionalChild--
				}
			}

			if insideStr {
				final += string(current)
			}

			continue
		} else {
			// check if it's our first child or not.
			if !firstChild {
				firstChild = true
				continue
			}

			// check if it's our first child or not.
			if !secondChild {
				secondChild = true
				continue
			}

			// check if it's our first child or not.
			if !thirdChild {
				thirdChild = true
				continue
			}

			// we have at most 3 official child.
			// we need only these childs.
			// but in the data, we will have more than these.
			// we don't know how many, maybe 4, maybe 5 or even 6.
			// which is why we have to use counter for them.
			additionalChild++
		}
	}

	return allStrs
}
