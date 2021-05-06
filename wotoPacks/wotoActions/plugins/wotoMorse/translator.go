// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrong"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// TranslateToMorse will translate the specified text value
// to the morse value
func TranslateToMorse(text string) string {
	return ToMorse(TranslateToBinary(text))
}

// TranslateToMorse will translate the specified text value
// to the morse value
func TranslateToBinary(text string) string {
	morseInit()
	var morseSequence string

	text = strings.ToUpper(text)
	text = strings.Trim(text, wv.SPACE_VALUE)
	textSlice := strings.Split(text, wv.EMPTY)

	for _, char := range textSlice {
		current := _alphabet[char]
		// check if the current character does exist
		// in our alphabet map or not.
		if wotoStrings.IsEmpty(&current) {
			// so, it seems it doesn't exist,
			// but let's do a little dirty trick here.
			for i, myChar := range char {
				if i != wv.BaseIndex {
					break
				}
				morseSequence += getSBinary(myChar) + wv.SPACE_VALUE
			}

		} else {
			morseSequence += _alphabet[char] + wv.SPACE_VALUE
		}
	}

	return morseSequence
}

// TranslateFromMorse will translate the morseSequence value to
// the morse value.
func TranslateFromMorse(morseSequence string) string {
	morseInit()
	var trs string // translated string

	morseSequence = ToBinary(strings.Trim(morseSequence, wv.SPACE_VALUE))
	morseSlice := strings.Split(morseSequence, wv.SPACE_VALUE)

	for _, char := range morseSlice {
		current := _reverseAlphabet[char]
		// check if the current character does exist
		// in our alphabet map or not.
		if wotoStrings.IsEmpty(&current) {
			// so, it seems it doesn't exist,
			// but let's do a little dirty trick here.
			trs += getSMorse(char)
		} else {
			trs += _reverseAlphabet[char]
		}
	}

	return trs
}

func getSBinary(r rune) string {
	str := strconv.FormatInt(int64(r), wv.BaseTwoIndex)
	off := strconv.FormatInt(wotoStrong.StrongOffSet, wv.BaseTwoIndex)
	return off + wv.UNDER + str
}

func getSMorse(binary string) string {
	if wotoStrings.IsEmpty(&binary) {
		return wv.EMPTY
	}
	// separate the lang code and the original char
	myStrings := strings.Split(binary, wv.UNDER)
	// ensure that lang code and original character exist
	if len(myStrings) != wv.BaseTwoIndex {
		return wv.EMPTY
	}
	// we should ensure that the lang code is exatcly the
	// same as our offset.
	num, err := strconv.ParseInt(myStrings[wv.BaseIndex],
		wv.BaseTwoIndex, wv.Base8Bit)
	if err != nil || num != wotoStrong.StrongOffSet {
		return wv.EMPTY
	}

	// now that everything is okay, we can convert the original
	// value.
	num, err = strconv.ParseInt(myStrings[wv.BaseOneIndex],
		wv.BaseTwoIndex, wv.Base64Bit)
	appSettings.GetExisting().SendSudo("HERE " + binary)
	if err != nil {
		return wv.EMPTY
	}

	return string(rune(num))
}
