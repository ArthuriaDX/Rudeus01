package wotoTranslate

import wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

// detect language constants keys and values.
const (
	//-----------------------------------------------------
	dHostUrl   = "https://detectlanguage.com/demo"
	gnuHostUrl = "https://api.gnuweeb.org/google_translate.php?" +
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
)

// google transtlate constants keys and values.
const (
	//-----------------------------------------------------

	requestType = "POST" // not optional
	gHostUrl    = "https://translate.google.com/_/" +
		"TranslateWebserverUi/data/" +
		"batchexecute?rpcids=MkEWBc" +
		"&f.sid=-6960075458768589634" +
		"&bl=boq_translate-webserver_20210512.09_p0" +
		"&hl=en-US&soc-app=1&soc-platform=1" +
		"&soc-device=1&_reqid=1662330&rt=c" // not optional

	//-----------------------------------------------------

	userAgentGKey   = "User-Agent" // not optional
	userAgentGValue = "Mozilla/5.0 " +
		"(X11; Ubuntu; Linux x86_64; rv:88.0) " +
		"Gecko/20100101 Firefox/88.0" // not optional

	//-----------------------------------------------------

	acceptGKey   = "Accept" // not optional
	acceptGValue = "*/*"    // not optional

	//-----------------------------------------------------
	acceptLanguageGKey   = "Accept-Language" // not optional
	acceptLanguageGValue = "en-US,en;q=0.5"  // not optional
	//-----------------------------------------------------
	refererGKey   = "Referer"                       // not optional
	refererGValue = "https://translate.google.com/" // not optional
	//-----------------------------------------------------
	xSameDomainGKey   = "X-Same-Domain" // not optional
	xSameDomainGValue = "1"             // not optional
	//-----------------------------------------------------
	xGoogBatchExecuteBgrGKey   = "X-Goog-BatchExecute-Bgr" // not optional
	xGoogBatchExecuteBgrGValue = "[\"!uLulu_bNAAZ-n43Xfp9" +
		"ChR5KyniU7RY7ACkAIwj8Rr5ZYBnKqvI3yOFfAcxDZqjGlJRAj7Wy" +
		"DbSIHd2rmrWJCLwm1AIAAAJAUgAAADJoAQcKALcwYzzTEVUDx7SCQT" +
		"1CoAv4iYyjg9RCfDFMYximkyqYTe38REBHCCV4VZFVBaph5E5VOlJm6Y" +
		"Lr8a_iniF72JIbG931cR_N2whV6a0OOTIxlYuY29VzYgH0lUipEtHoT7O0" +
		"BcWxFMu-mhiHCgIf5CDRGVFl4Y6GWWBXgNqOv2LMtr8nziYtayYIrEvFpV" +
		"wFEITAZp4-3QT3MYAdJ2U7wrHsW8eWZqQzmOC2biGnKa_YYd-NIbvIi_CZAc" +
		"9WmbBfsDWF2NLtoug2oUVk4oyRZoXMnRtRpsuXOp1ydcgogFXl1PKNDvctRzp" +
		"A4E1dJWjLskR1Ht7HpCO611v9o6BePdBLB6-rM-jQOejGLiJvqq-vS3rpCSr" +
		"TRR8OkyZh0emPZTP6B4dcOz_KH_0IYQghx2LnAxy5eaA5DDzYAECp-TsCb-" +
		"AvbLgRVA-PkqoargQ99NyBlxv9CZQngEtbhwyXzSxpdFCPhikJUIPwUPMN5Gc" +
		"1Y5B5HTNh9xnYndYneSQWtXRUHFNW1nMOCenMnoHEN0iq8U_OiYHnakZPlm" +
		"EG752mnidgLBT2CJVLkbTPVUoMN7HiFUTk-koWIzhAOdSWznIHanHiQr" +
		"20OmXSB5uURXCm-3_uNHR25vSJnDw3-MbEKdMmtlMDcyzU8sfwZ-ilCj" +
		"FAby6hJpBJq1MAVibjxniaed6z38EiLqdnCR_vJVhnXZ01cb2Ua9QuvY5" +
		"WtEhvpXGmJ6K9KdppK9n9VOQP9g2QXzfem3WIahR7a0AN_98Gtv_" +
		"Df5tWUfMpyj6hSwSd_8ZdnLkTv5VNPy-R0eWmPIwrmQq00IzvGs42" +
		"VYhrkPNvJhG4FuRdvebmsu63yRHjGX6zt9U_EJOehdii\"" +
		",null,null,51,null,null,null,0]" // not optional
	//-----------------------------------------------------
	contentTypeGKey   = "Content-Type" // not optional
	contentTypeGValue = "application/x-www-form-urlencoded;" +
		"charset=utf-8" // not optional
	//-----------------------------------------------------
	originGKey   = "Origin"                       // not optional
	originGValue = "https://translate.google.com" // not optional
	//-----------------------------------------------------
	gDNTGKey   = "DNT" // not optional
	gDNTGValue = "1"   // not optional
	//-----------------------------------------------------
	connectionGKey   = "Connection" // not optional
	connectionGValue = "keep-alive" // not optional
	//-----------------------------------------------------
	cookieGKey   = "Cookie" // not optional
	cookieGValue = "NID=215=loqde-FVPxqRXGmigy_" +
		"yQSS8T9ElYCYG3hZYItUEtpIP2aoDaKQPVWLEt5x9" +
		"jlB3Rju7OOztJe1vN4TCYuYz996AdeYk3GuGiXFBg" +
		"B3jJ1y3bvWzTi3EkSXfsAjXgaI0IY4o56gTgT20QOa" +
		"I-8akMGD40bZeaN5LTnPEaAksP98; OTZ=5978478_56_56__56_" // not optional
	//-----------------------------------------------------
	fReqGKey = "f.req=" // not optional
	// GValue1 + (TEXT(urlencoded)) + GValue2 + (auto) +
	// GValue3 + (en) + GValue4
	fReqGValue1 = fReqGKey + "%5B%5B%5B%22MkEWBc%22%2C%22%5B%5B%5C%22"
	fReqGValue2 = "%5C%22%2C%5C%22"
	fReqGValue3 = "%5C%22%2C%5C%22"
	fReqGValue4 = "%5C%22%2Ctrue%5D%2C%5Bnull" +
		"%5D%5D%22%2Cnull%2C%22generic%22%5D%5D%5D&"
)

const (
	DoubleQS       = wv.BackSlash + wv.DoubleQ
	DoubleQSP      = DoubleQS + wv.CAMA
	NonEscapeN     = "\\n"
	NonEscapeNV    = NonEscapeN + wv.CAMA
	HttpRm         = "af.httprm"
	E4Value        = "\"e\"" + wv.CAMA + "4" + wv.CAMA
	NullCValue     = wv.NullStr + wv.CAMA
	GenericStr     = wv.CAMA + "\"generic\""
	NullCValueR    = wv.CAMA + wv.NullStr
	NeQ            = "\\n\""
	NullN          = "\n\"" + wv.CAMA + "null"
	DiValue        = "\"di\""
	AkCloseQ       = "}'"
	WrbFr          = "\"wrb.fr\",\"MkEWBc\""
	BoldOpen       = "\\u003cb\\u003e"
	BoldClose      = "\\u003c/b\\u003e"
	MiddleWave     = DoubleQSP + DoubleQS
	WrongNessOpen  = "\\u003cb\\\\u003e\\\\u003ci\\\\u003e"
	WrongNessClose = "\\\\u003c/i\\\\u003e\\\\u003c/b\\\\u003e"
)

const (
	errTrGnuText = "Couldn't translate the text using GNU/WEEB API >-<!"
)
