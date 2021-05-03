// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// TranslateToMorse will translate the specified text value
// to the morse value (only english and numbers are supported for now)
func TranslateToMorse(text string) string {
	morseInit()
	var morseSequence string

	text = strings.ToUpper(text)
	text = strings.Trim(text, wotoValues.SPACE_VALUE)
	textSlice := strings.Split(text, wotoSecurity.EMPTY)

	for _, char := range textSlice {
		morseSequence += _alphabet[char] + wotoValues.SPACE_VALUE
	}

	return morseSequence
}

// TranslateFromMorse will translate the morseSequence value to
// the morse value.
func TranslateFromMorse(morseSequence string) string {
	morseInit()
	var trs string // translated string

	morseSequence = strings.Trim(morseSequence, wotoValues.SPACE_VALUE)
	morseSlice := strings.Split(morseSequence, wotoValues.SPACE_VALUE)

	for _, char := range morseSlice {
		trs += _reverseAlphabet[char]
	}

	return trs
}
