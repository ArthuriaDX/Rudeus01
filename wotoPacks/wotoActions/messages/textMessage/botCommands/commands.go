package botCommands

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// IsCommand will check if the _text is a command or not.
func IsCommand(_text *string) bool {
	return isCommand1(_text) || isCommand2(_text)
}

// isCommand1 will check if the _text
// has first prefex or not.
func isCommand1(_text *string) bool {
	return strings.HasPrefix(*_text, wotoValues.COMMAND_PREFEX1)
}

// isCommand2 will check if the _text
// has first prefex or not.
func isCommand2(_text *string) bool {
	return strings.HasPrefix(*_text, wotoValues.COMMAND_PREFEX2)
}

func HandleCommand(_message *tgbotapi.Message) {

}

//func main() {
//    input := `xxxxx:yyyyy:zzz.aaa.bbb.cc:dd:ee:ff`
//    a := strings.FieldsFunc(input, Split)
//    t := Target{a[0], a[1], a[2], a[3], a[4], a[5], a[6]}
//    fmt.Println(t) // {xxxxx yyyyy zzz aaa bbb cc dd}
//}
//func Split(r rune) bool {
//    return r == ':' || r == '.'
//}
