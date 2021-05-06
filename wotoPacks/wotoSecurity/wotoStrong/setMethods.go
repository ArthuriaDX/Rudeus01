// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

// _setValue will set the bytes value of the StrongString.
func (_s *StrongString) _setValue(theString string) {
	_s._setValueByBytes(convertToBytes(theString))
}

// _setValueByBytes will set the bytes value directly.
func (_s *StrongString) _setValueByBytes(_b []rune) {
	_s._value = _b
}
