package wotoUD

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
	nextText     = "next page"
	previousText = "previous page"
)
