// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import (
	"strconv"
	"strings"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	MaxUIDIndex    = 9
	BaseUIDIndex   = 1
	UIDIndexOffSet = 1
)

const (
	// the StrongBase is Hexadecimal, which is base 16.
	//StrongBase = 0x010
	// the StrongOffSet is the offset value of each character in the
	// strong string.
	StrongOffSet = 0x00D
)

// the StrongString used in the program for High-security!
type StrongString struct {
	_value []rune
}

// generateStrongString will generate a new StrongString
// with the specified non-encoded string value.
func GenerateStrongString(_s string) StrongString {
	_strong := StrongString{}
	_strong._setValue(_s)
	return _strong
}

// getStrong will give you an already encoded StrongString
// which provided by method _getString() method and saved in the
// data base.
func GetStrong(_s string) StrongString {
	_strong := StrongString{}
	myBytes := toByteArrayS(_s, wv.LineStr)
	_strong._setValueByBytes(myBytes)
	return _strong
}

// convert the specified string array (encoded) to byte array.
func _toByteArray(myStrings []string) []rune {
	myBytes := make([]rune, 0)
	for _, _current := range myStrings {
		_myInt, _ := strconv.Atoi(_current)
		myBytes = append(myBytes, rune(_myInt+StrongOffSet))
	}
	return myBytes
}

// convert the specified string (encoded) to byte array.
func toByteArrayS(theString string, separator string) []rune {
	return _toByteArray(ws.FixSplit(strings.Split(theString, separator)))
}

// convertToBytes will convert an ordinaryString (non-encoded) to byte array.
func convertToBytes(ordinaryString string) []rune {
	_runes := []rune(ordinaryString)
	finalRune := make([]rune, len(_runes), cap(_runes))
	for _i, _current := range _runes {
		finalRune[_i] = _current + StrongOffSet
	}
	return finalRune
}

// ConvertToString will convert the specified byte arrays to string.
func ConvertToString(_b []rune) string {
	_total := wv.EMPTY
	for _, _current := range _b {
		_total += string(_current - StrongOffSet)
	}
	return _total
}
