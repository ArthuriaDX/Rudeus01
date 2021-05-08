// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package tgMessages

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgForbidden"
)

func GetTgError(err error) interfaces.TgError {
	str := err.Error()
	if strings.Contains(str, tgForbidden.FORBIDDEN) {
		fr, ok := tgForbidden.GetForbidden(err)
		if !ok {
			return nil
		}
		return fr
	}
	return nil
}
