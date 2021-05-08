// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrings

import (
	"strings"

	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func Split(_s string, separator ...string) []string {
	if len(separator) == wv.BaseIndex {
		return nil
	}
	_firstSet := false
	var (
		_shallow   string
		_myStrings []string
	)
	for _, _current := range separator {
		if strings.Contains(_s, _current) {
			if !_firstSet {
				_myStrings = strings.Split(_s, _current)
				_myStrings = FixSplit(_myStrings)
				_firstSet = true
				continue
			}
			_shallow = strings.Join(_myStrings, wv.EMPTY)
			_myStrings = strings.Split(_shallow, _current)
			_myStrings = FixSplit(_myStrings)
		}
	}
	if !_firstSet {
		_myStrings = make([]string, wv.BaseIndex)
		_myStrings[wv.BaseIndex] = _s
	}
	return _myStrings
}

// FixSplit will fix the bullshit bug in the
// Split function (which is not ignoring the spaces between strings).
func FixSplit(_myStrings []string) []string {
	final := make([]string, 0, cap(_myStrings))
	for _, _current := range _myStrings {
		if !IsEmpty(&_current) {
			final = append(final, _current)
		}
	}
	return final
}

// IsEmpty function will check if the passed-by
// string value is empty or not.
func IsEmpty(_s *string) bool {
	return _s == nil || *_s == wv.EMPTY
}
