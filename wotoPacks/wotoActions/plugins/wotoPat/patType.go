// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoPat

// the pat type, when sending a files, we should use this type.
type patType uint8

// the pat types values.
const (
	NONE      patType = 0
	PHOTO_PAT patType = 1
)
