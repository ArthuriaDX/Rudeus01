package wotoMD

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

type wotoMarkDown struct {
	_value string
}

func GetNormal(value string) interfaces.WMarkDown {
	final := repairValue(value)
	return toWotoMD(final)
}

func GetBold(value string) interfaces.WMarkDown {
	final := repairValue(value)
	final = string(CHAR_S3) + final + string(CHAR_S3)
	return toWotoMD(final)
}

func GetItalic(value string) interfaces.WMarkDown {
	final := repairValue(value)
	final = string(CHAR_S4) + final + string(CHAR_S4)
	return toWotoMD(final)
}

func GetMono(value string) interfaces.WMarkDown {
	final := repairValue(value)
	final = string(CHAR_S16) + final + string(CHAR_S16)
	return toWotoMD(final)
}

func IsSpecial(r rune) bool {
	for _, current := range _sChars {
		if r == current {
			return true
		}
	}
	return false
}

func toWotoMD(value string) interfaces.WMarkDown {
	if wotoStrings.IsEmpty(&value) {
		return nil
	}
	myMD := wotoMarkDown{
		_value: value,
	}
	return &myMD
}

func repairValue(value string) string {
	finally := wv.EMPTY
	escape := false
	lasEscape := false
	escapeCount := wv.BaseIndex
	for i, current := range value {
		escape = (current == CHAR_S1)
		if IsSpecial(current) {
			if escape {
				escapeCount++
			} else {
				escapeCount = wv.BaseIndex
			}
			if i != wv.BaseIndex {
				if !lasEscape {
					finally += string(CHAR_S1) + string(current)
				} else {
					finally += string(current)
				}
			} else {
				finally += string(CHAR_S1) + string(current)
			}
		} else {
			if escapeCount != wv.BaseIndex {
				tmpR := escapeCount % wv.BaseTwoIndex
				if tmpR != wv.BaseIndex {
					finally += string(CHAR_S1) + string(current)
				} else {
					finally += string(current)
				}
			} else {
				finally += string(current)
			}
		}
		lasEscape = escape
	}
	return finally
}
