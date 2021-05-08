// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import (
	"reflect"
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// GetValue will give you the real value of this StrongString.
func (_s *StrongString) GetValue() *string {
	realString := ConvertToString(_s._value)
	return &realString
}

// length method, will give you the length-as-int of this StrongString.
func (_s *StrongString) Length() int {
	return len(*_s.GetValue())
}

// isEmpty will check if this StrongString is empty or not.
func (_s *StrongString) IsEmpty() bool {
	return ws.IsEmpty(_s.GetValue())
}

// isEqual will check if the passed-by-value in the arg is equal to this
// StrongString or not.
func (_s *StrongString) IsEqual(_q interfaces.QString) bool {
	if reflect.TypeOf(_q) != reflect.TypeOf(_s) {
		return *_q.GetValue() == *_s.GetValue()
	}
	_strong, _ok := _q.(*StrongString)
	if !_ok {
		return false
	}
	// check if the length of them are equal or not.
	if len(_s._value) != len(_strong._value) {
		//fmt.Println(len(_s._value), len(_strong._value))
		return false
	}
	for i := 0; i < len(_s._value); i++ {
		if _s._value[i] != _strong._value[i] {
			//fmt.Println(_s._value[i], _strong._value[i])
			return false
		}
	}
	return true
}

// GetString will give you an encoded string with the High-security
// level which you should use it in the database.
func (_s *StrongString) GetString() string {
	var _current int
	var total = wv.EMPTY
	for _, b := range _s._value {
		_current = int(b)
		total += strconv.Itoa(_current)
		total += wv.LineStr
	}
	return total
}

// GetIndexV method will give you the rune in _index.
func (_s *StrongString) GetIndexV(_index int) rune {
	l := len(_s._value)
	if l >= len(_s._value) || l < wv.BaseIndex {
		return _s._value[wv.BaseIndex]
	}
	return _s._value[_index]
}
