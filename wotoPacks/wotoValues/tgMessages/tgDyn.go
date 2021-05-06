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
