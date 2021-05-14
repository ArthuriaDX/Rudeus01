// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

type QString interface {
	Length() int
	IsEmpty() bool
	GetValue() string
	GetIndexV(_index int) rune
	IsEqual(_strong QString) bool
}
