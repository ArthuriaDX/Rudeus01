// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMD

import (
	"reflect"

	infc "github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func (_m *wotoMarkDown) Append(v infc.WMarkDown) infc.WMarkDown {
	if reflect.TypeOf(v) == reflect.TypeOf(_m) {
		md := v.(*wotoMarkDown)
		str := _m._value + wotoValues.SPACE_VALUE + md._value
		wmd := wotoMarkDown{
			_value: str,
		}
		return &wmd
	}
	return nil
}

func (_m *wotoMarkDown) ToString() string {
	return _m._value
}
