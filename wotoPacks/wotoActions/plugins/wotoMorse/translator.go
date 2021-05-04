// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
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
	textSlice := strings.Split(text, wotoSecurity.EMPTY)

	for _, char := range textSlice {
		morseSequence += _alphabet[char] + wv.SPACE_VALUE
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
		trs += _reverseAlphabet[char]
	}

	return trs
}
