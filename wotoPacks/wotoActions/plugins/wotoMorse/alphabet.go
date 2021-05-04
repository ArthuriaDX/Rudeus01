// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

import (
	"strings"

	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// _alphabet is a mapping of Alpha numeric characters to Morse code
var _alphabet map[string]string

var _reverseAlphabet map[string]string

func morseInit() {
	if _alphabet == nil {
		init_alphabet()
	}
	if _reverseAlphabet == nil {
		init_reverse()
	}
}
func init_alphabet() {
	setMap()
}

func init_reverse() {
	if _alphabet == nil {
		return
	}
	if _reverseAlphabet == nil {
		_reverseAlphabet = make(map[string]string)
	}
	for k, v := range _alphabet {
		_reverseAlphabet[v] = k
	}
}

// ToBinary will convert a morse string to binary string.
// for example it will convert ".-.-" to "0101"
func ToBinary(value string) string {
	str := strings.ReplaceAll(value, wv.DotStr, wv.BaseIndexStr)
	str = strings.ReplaceAll(str, wv.LineStr, wv.BaseOneIndexStr)

	return str
}

// ToMorse will convert a binary string to morse string.
// for example it will convert "0101" to ".-.-"
func ToMorse(value string) string {
	str := strings.ReplaceAll(value, wv.BaseIndexStr, wv.DotStr)
	str = strings.ReplaceAll(str, wv.BaseOneIndexStr, wv.LineStr)

	return str
}
