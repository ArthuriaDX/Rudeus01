// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

const (
	MaxUIDIndex    = 9
	BaseUIDIndex   = 1
	UIDIndexOffSet = 1
)

const (
	StrongOffSet = 13
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
