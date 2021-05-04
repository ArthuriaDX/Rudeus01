// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package pTools

import (
	"strings"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func IsFlag(value string) bool {
	//log.Println("In IsFlag, I'm going to check", value)
	return strings.HasPrefix(value, wv.FLAG_PREFEX)
}

// GetFlags will give you all of the flags of this Arg value.
func (_a *Arg) GetFlags() []string {
	cap := len(*_a) - wv.BaseOneIndex
	strs := make([]string, wv.BaseIndex, cap)
	var current string
	for _, str := range *_a {
		//log.Println("in range, we are going to check ", str)
		//appSettings.GetExisting().SendSudo("Is NIL!") // TODO
		if IsFlag(str) {
			//log.Println("I've already checked str and it was flag.")
			current = strings.TrimPrefix(str, wv.FLAG_PREFEX)
			current = strings.ToLower(current)
			strs = append(strs, current)
		}
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
		//appSettings.GetExisting().SendSudo("Is NIL!")
		return false
	}
	//appSettings.GetExisting().SendSudo(fmt.Sprint(len(flags)))
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
	//log.Println("After Trimming, the flag is:", flag)
	flags := _a.GetFlags()

	//appSettings.GetExisting().SendSudo(fmt.Sprint("a.getFlags result:", flags))
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
