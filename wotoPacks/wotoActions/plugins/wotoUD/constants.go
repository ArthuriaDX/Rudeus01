package wotoUD

import wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

// the url for urban dictionary public API.
const (
	hostUrl = "https://api.urbandictionary.com/v0/define?term="
)

// error messages for wotoUD package.
const (
	emptyWordErr = "wotoUD: word cannot be empty"
	nilRespErr   = "wotoUD: an unexpected error, response or its body is nil"
	jDataErr     = "wotoUD: json data was nil or zero"
)

const (
	nextText = "next page"     // the next button Text
	preText  = "previous page" // the previous button Text
	footnote = "\n\n-----------------------\nPage " +
		wv.FORMAT_VALUE + " of " + wv.FORMAT_VALUE
	listNAvailble = "list is not available anymore."
	notAllowed    = "Owo ~ You are not allowed to use this button!"
)
