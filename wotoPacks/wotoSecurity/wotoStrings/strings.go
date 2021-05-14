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

	final := _s
	for _, myStr := range separator {
		final = strings.ReplaceAll(final, myStr, sepStr)
	}
	return FixSplit(strings.Split(final, sepStr))
}

func SplitSlice(_s string, separator []string) []string {
	if len(separator) == wv.BaseIndex {
		return nil
	}

	final := _s
	for _, myStr := range separator {
		final = strings.ReplaceAll(final, myStr, sepStr)
	}

	return FixSplit(strings.Split(final, sepStr))
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
