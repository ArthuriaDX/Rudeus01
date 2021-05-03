// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package pTools

import (
	"strings"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func IsFlag(value string) bool {
	return strings.HasPrefix(value, wv.FLAG_PREFEX)
}

// GetFlags will give you all of the flags of this Arg value.
func (_a *Arg) GetFlags() []string {
	cap := len(*_a) - wv.BaseOneIndex
	strs := make([]string, wv.BaseIndex, cap)
	for _, str := range *_a {
		if IsFlag(str) {
			strs = append(strs, strings.TrimPrefix(str, wv.FLAG_PREFEX))
		}
	}
	// prevents from returning an empty array!
	if len(strs) == wv.BaseOneIndex {
		return nil
	}
	return strs
}

// HasFlag will check if at least one of the flags is present in
// flags or not.
// notice: this method, will check for at least one of them,
// it means if it finds one of the flags value, it will return true.
// if you want to search for all of them, consider using HasFlags.
func (_a *Arg) HasFlag(flags ...string) bool {
	if flags == nil {
		return false
	}
	for _, flag := range flags {
		if _a.hasFlag(flag) {
			return true
		}
	}
	return false
}

func (_a *Arg) hasFlag(flag string) bool {
	if ws.IsEmpty(&flag) {
		return false
	}
	// don't need to check if it has this prefex or not!
	// right, I don't trust myself!
	// also TrimPrefix function has a checker in itself.
	flag = strings.TrimPrefix(flag, wv.FLAG_PREFEX)
	flags := _a.GetFlags()
	if flags == nil {
		return false
	}
	for _, mflag := range flags {
		if mflag == flag {
			return true
		}
	}
	return false
}

// HasFlags will check if all of the flags are present in
// flags or not.
// notice: this method, will check for all of them,
// it means if it doesn't find one of the flags value, it will return false.
// if you want to search for at least one of them, consider using HasFlag.
func (_a *Arg) HasFlags(flags ...string) bool {
	if flags == nil {
		return false
	}
	for _, flag := range flags {
		if !_a.hasFlag(flag) {
			return false
		}
	}
	return true
}
