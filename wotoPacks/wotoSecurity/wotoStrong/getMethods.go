// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import (
	"reflect"
	"strings"

	tf "github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// GetValue will give you the real value of this StrongString.
func (_s *StrongString) GetValue() string {
	return string(_s._value)
}

// length method, will give you the length-as-int of this StrongString.
func (_s *StrongString) Length() int {
	return len(_s._value)
}

// isEmpty will check if this StrongString is empty or not.
func (_s *StrongString) IsEmpty() bool {
	return _s._value == nil || len(_s._value) == wv.BaseIndex
}

// isEqual will check if the passed-by-value in the arg is equal to this
// StrongString or not.
func (_s *StrongString) IsEqual(_q tf.QString) bool {
	if reflect.TypeOf(_q) != reflect.TypeOf(_s) {
		return _q.GetValue() == _s.GetValue()
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

// GetIndexV method will give you the rune in _index.
func (_s *StrongString) GetIndexV(_index int) rune {
	if _s.IsEmpty() {
		return wv.BaseIndex
	}

	l := len(_s._value)

	if _index >= l || l < wv.BaseIndex {

		return _s._value[wv.BaseIndex]
	}

	return _s._value[_index]
}

// HasSuffix will check if at least there is one suffix is
// presents in this StrongString not.
// the StrongString should ends with at least one of these suffixes.
func (_s *StrongString) HasSuffix(values ...string) bool {
	for _, s := range values {
		if strings.HasSuffix(_s.GetValue(), s) {
			return true
		}
	}

	return false
}

// HasSuffixes will check if all of the suffixes are
// present in this StrongString or not.
// the StrongString should ends with all of these suffixes.
// usage of this method is not recommended, since you can use
// HasSuffix method with only one string (the longest string).
// this way you will just use too much cpu resources.
func (_s *StrongString) HasSuffixes(values ...string) bool {
	for _, s := range values {
		if !strings.HasSuffix(_s.GetValue(), s) {
			return false
		}
	}

	return true
}

// HasPrefix will check if at least there is one prefix is
// presents in this StrongString or not.
// the StrongString should starts with at least one of these prefixes.
func (_s *StrongString) HasPrefix(values ...string) bool {
	for _, s := range values {
		if strings.HasPrefix(_s.GetValue(), s) {
			return true
		}
	}

	return false
}

// HasPrefixes will check if all of the prefixes are
// present in this StrongString or not.
// the StrongString should ends with all of these suffixes.
// usage of this method is not recommended, since you can use
// HasSuffix method with only one string (the longest string).
// this way you will just use too much cpu resources.
func (_s *StrongString) HasPrefixes(values ...string) bool {
	for _, s := range values {
		if !strings.HasPrefix(_s.GetValue(), s) {
			return false
		}
	}

	return true
}

func (_s *StrongString) Split(qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSlice(_s.GetValue(), ToStrSlice(qs))
	return ToQSlice(strs)
}

func (_s *StrongString) SplitStr(qs ...string) []tf.QString {
	strs := wotoStrings.SplitSlice(_s.GetValue(), qs)
	return ToQSlice(strs)
}

func (_s *StrongString) ToQString() tf.QString {
	return _s
}
