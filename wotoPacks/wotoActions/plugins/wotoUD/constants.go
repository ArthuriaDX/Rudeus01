package wotoUD

import wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

//---------------------------------------------------------

// the url for urban dictionary public API.
const (
	hostUrl      = "https://api.urbandictionary.com/v0/define?term="
	googleSimUrl = "https://www.google.com/complete/" +
		"search?q=" + wv.DY_WOTO_TEXT +
		"&cp=0&client=gws-wiz" +
		"&xssi=t&gs_ri=gws-wiz&hl=en-US&authuser=0&" +
		"pq=" + wv.DY_WOTO_TEXT +
		"&psi=qUycYJ_XLseNkwXX7oGgBg.1620855979375&nolsbt=1&dpr=1"
)

//---------------------------------------------------------
// error messages for wotoUD package.
const (
	emptyWordErr = "wotoUD: word cannot be empty"
	emptyUrlErr  = "wotoUD: returned url was empty"
	correctErr   = "wotoUD: correct word not found in google's database"
	nilRespErr   = "wotoUD: an unexpected error response or its body is nil"
	jDataErr     = "wotoUD: json data was nil or zero"
)

//---------------------------------------------------------

// the constants used for interact with users.
const (
	nextText = "next page"     // the next button Text
	preText  = "previous page" // the previous button Text
	footnote = "\n\n-----------------------\nPage " +
		wv.FORMAT_VALUE + " of " + wv.FORMAT_VALUE
	listNAvailble = "list is not available anymore."
	notAllowed    = "Owo ~ You are not allowed to use this button!"
	notFound      = "couldn't find \"" +
		wv.FORMAT_VALUE + "\"\n"
	simNotice = "hmmmm, here are some similarities: \n"
	numSep    = " - "
)

//---------------------------------------------------------
