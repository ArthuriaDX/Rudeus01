package wotoTranslate

import wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

const (
	//-----------------------------------------------------
	dHostUrl = "https://detectlanguage.com/demo"
	gHostUrl = "https://api.gnuweeb.org/google_translate.php?" +
		"fr=" + wv.FORMAT_VALUE +
		"&to=" + wv.FORMAT_VALUE +
		"&text=" + wv.FORMAT_VALUE // fr, to, text
	//-----------------------------------------------------
	userAgentKey   = "User-Agent"
	userAgentValue = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0) " +
		"Gecko/20100101 Firefox/88.0"
	//-----------------------------------------------------
	acceptKey   = "Accept"
	acceptValue = "*/*"
	//-----------------------------------------------------
	acceptLanguageKey   = "Accept-Language"
	acceptLanguageValue = "en-US,en;q=0.5"
	//-----------------------------------------------------
	refererKey   = "Referer"
	refererValue = "https://detectlanguage.com/"
	//-----------------------------------------------------
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json"
	//-----------------------------------------------------
	originKey   = "Origin"
	originValue = "https://detectlanguage.com"
	//-----------------------------------------------------
	connectionKey   = "Connection"
	connectionValue = "keep-alive"
	//-----------------------------------------------------
	cookieKey   = "Cookie"
	cookieValue = "__cfduid=d444ac98f951ec573ac1e2f7631908ceb1620422371;" +
		" _ga=GA1.2.1743024693.1620422382; " +
		"_detectlanguage_web_session=d79%2FR6%2BPq1%2FdkpJ2OWCsLoCaU6ru2Xsfn" +
		"ahO2%2BLe4V70lrWWZ6mj45u6Q2zh%2BagfEbLQQhFlOqk0z5OtXDHpFzQh8unIDlFm" +
		"OftjVnLpzUE3OO9nKybfwhttLeunN5YTOQWWATqyNj7ouqaqNwMG2xiU5gq89TAQUQB" +
		"6rEvNMjooFggSG3LsJO9bm9OCpIJRjAIaRkp2%2FBRlWPbh1eoGO8MwyP7FBVRTrwWF" +
		"oqYJW1zERRBSAJJc6B62ITseifLLKgh5zqUSP3%2BDgODMy7roqF5Quhm31TPFrSRUv" +
		"REywIGhArg%3D--GJmuqPbuSxkfw64o--oU%2B7kGsa0GW6ThldCik97A%3D%3D;" +
		" _gid=GA1.2.2108164705.1620509532; _gat_gtag_UA_133799_14=1"
	//-----------------------------------------------------
	teKey   = "TE"
	teValue = "Trailers"
	//-----------------------------------------------------
	qKey = "q" // qValue is the text.
	//-----------------------------------------------------
	preLeft  = "<pre>"
	preRight = "</pre>"
	//-----------------------------------------------------
)
