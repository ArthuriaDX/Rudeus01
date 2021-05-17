package wotoLang

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

type Language string

var langList map[string]string
var langListR map[string]string

// languages constants
const (
	L_af    = "af"    // Afrikaans
	L_am    = "am"    // Amharic
	L_ar    = "ar"    // Arabic
	L_az    = "az"    // Azerbaijani
	L_be    = "be"    // Belarusian
	L_bg    = "bg"    // Bulgarian
	L_bn    = "bn"    // Bengali
	L_bs    = "bs"    // Bosnian
	L_ca    = "ca"    // Catalan
	L_ceb   = "ceb"   // Chechen
	L_co    = "co"    // Corsican
	L_cs    = "cs"    // Czech
	L_cy    = "cy"    // Welsh
	L_da    = "da"    // Danish
	L_de    = "de"    // German
	L_el    = "el"    // Greek
	L_en    = "en"    // English
	L_eo    = "eo"    // Esperanto
	L_es    = "es"    // Spanish
	L_et    = "et"    // Estonian
	L_eu    = "eu"    // Basque
	L_fa    = "fa"    // Persian
	L_fi    = "fi"    // Finnish
	L_fr    = "fr"    // French
	L_fy    = "fy"    // WesternFrisian
	L_ga    = "ga"    // Irish
	L_gd    = "gd"    // Gaelic
	L_gl    = "gl"    // Galician
	L_gu    = "gu"    // Gujarati
	L_ha    = "ha"    // Hausa
	L_haw   = "haw"   // ???
	L_hi    = "hi"    // Hindi
	L_hmn   = "hmn"   // ???
	L_hr    = "hr"    // Croatian
	L_ht    = "ht"    // Haitian
	L_hu    = "hu"    // Hungarian
	L_hy    = "hy"    // Armenian
	L_id    = "id"    // Indonesian
	L_ig    = "ig"    // Igbo
	L_is    = "is"    // Icelandic
	L_it    = "it"    // Italian
	L_iw    = "iw"    // Hebrew
	L_ja    = "ja"    // Japanese
	L_jw    = "jw"    // ???
	L_ka    = "ka"    // Georgian
	L_kk    = "kk"    // Kazakh
	L_km    = "km"    // Central Khmer
	L_kn    = "kn"    // Kannada
	L_ko    = "ko"    // Korean
	L_ku    = "ku"    // Kurdish
	L_ky    = "ky"    // Kirghiz
	L_la    = "la"    // Latin
	L_lb    = "lb"    // Luxembourgish
	L_lo    = "lo"    // Lao
	L_lt    = "lt"    // Lithuanian
	L_lv    = "lv"    // Latvian
	L_mg    = "mg"    // Malagasy
	L_mi    = "mi"    // Maori
	L_mk    = "mk"    // Macedonian
	L_ml    = "ml"    // Malayalam
	L_mn    = "mn"    // Mongolian
	L_mr    = "mr"    // Marathi
	L_ms    = "ms"    // Malay
	L_mt    = "mt"    // Maltese
	L_my    = "my"    // Burmese
	L_ne    = "ne"    // Nepali
	L_nl    = "nl"    // Dutch
	L_no    = "no"    // Norwegian
	L_ny    = "ny"    // Chichewa
	L_pa    = "pa"    // Punjabi
	L_pl    = "pl"    // Polish
	L_ps    = "ps"    // Pashto
	L_pt    = "pt"    // Portuguese
	L_ro    = "ro"    // Romanian
	L_ru    = "ru"    // Russian
	L_sd    = "sd"    // Sindhi
	L_si    = "si"    // Sinhala
	L_sk    = "sk"    // Slovak
	L_sl    = "sl"    // Slovenian
	L_sm    = "sm"    // Samoan
	L_sn    = "sn"    // Shona
	L_so    = "so"    // Somali
	L_sq    = "sq"    // Albanian
	L_sr    = "sr"    // Serbian
	L_st    = "st"    // Southern Sotho
	L_su    = "su"    // Sundanese
	L_sv    = "sv"    // Swedish
	L_sw    = "sw"    // Swahili
	L_ta    = "ta"    // Tamil
	L_te    = "te"    // Telugu
	L_tg    = "tg"    // Tajik
	L_th    = "th"    // Thai
	L_tl    = "tl"    // Tagalog
	L_tr    = "tr"    // Turkish
	L_uk    = "uk"    // Ukrainian
	L_ur    = "ur"    // Urdu
	L_uz    = "uz"    // Uzbek
	L_vi    = "vi"    // Vietnamese
	L_xh    = "xh"    // Xhosa
	L_yi    = "yi"    // Yiddish
	L_yo    = "yo"    // Yoruba
	L_zh    = "zh"    // Chinese
	L_zh_CN = "zh_CN" // Chinese
	L_zh_TW = "zh_TW" // Chinese
	L_zu    = "zu"    // Zulu
)

func initLang() {
	initLangMap()
	initLangReseve()
}

func initLangMap() {
	if langList == nil {
		langList = map[string]string{
			L_af:    "Afrikaans", // Afrikaans
			L_am:    "Amharic",   // Amharic
			L_ar:    "Arabic",
			L_az:    "Azerbaijani",
			L_be:    "Belarusian",
			L_bg:    "Bulgarian",
			L_bn:    "Bengali",
			L_bs:    "Bosnian",
			L_ca:    "Catalan",
			L_ceb:   "Chechen",
			L_co:    "Corsican",
			L_cs:    "Czech",
			L_cy:    "Welsh",
			L_da:    "Danish",
			L_de:    "German",
			L_el:    "Greek",
			L_en:    "English",
			L_eo:    "Esperanto",
			L_es:    "Spanish",
			L_et:    "Estonian",
			L_eu:    "Basque",
			L_fa:    "Persian",
			L_fi:    "Finnish",
			L_fr:    "French",
			L_fy:    "WesternFrisian",
			L_ga:    "Irish",
			L_gd:    "Gaelic",
			L_gl:    "Galician",
			L_gu:    "Gujarati",
			L_ha:    "Hausa",
			L_haw:   "haw", // ???
			L_hi:    "Hindi",
			L_hmn:   "hmn", // ???
			L_hr:    "Croatian",
			L_ht:    "Haitian",
			L_hu:    "Hungarian",
			L_hy:    "Armenian",
			L_id:    "Indonesian",
			L_ig:    "Igbo",
			L_is:    "Icelandic",
			L_it:    "Italian",
			L_iw:    "Hebrew",
			L_ja:    "Japanese",
			L_jw:    "jw", // ???
			L_ka:    "Georgian",
			L_kk:    "Kazakh",
			L_km:    "Central Khmer",
			L_kn:    "Kannada",
			L_ko:    "Korean",
			L_ku:    "Kurdish",
			L_ky:    "Kirghiz",
			L_la:    "Latin",
			L_lb:    "Luxembourgish",
			L_lo:    "Lao",
			L_lt:    "Lithuanian",
			L_lv:    "Latvian",
			L_mg:    "Malagasy",
			L_mi:    "Maori",
			L_mk:    "Macedonian",
			L_ml:    "Malayalam",
			L_mn:    "Mongolian",
			L_mr:    "Marathi",
			L_ms:    "Malay",
			L_mt:    "Maltese",
			L_my:    "Burmese",
			L_ne:    "Nepali",
			L_nl:    "Dutch",
			L_no:    "Norwegian",
			L_ny:    "Chichewa",
			L_pa:    "Punjabi",
			L_pl:    "Polish",
			L_ps:    "Pashto",
			L_pt:    "Portuguese",
			L_ro:    "Romanian",
			L_ru:    "Russian",
			L_sd:    "Sindhi",
			L_si:    "Sinhala",
			L_sk:    "Slovak",
			L_sl:    "Slovenian",
			L_sm:    "Samoan",
			L_sn:    "Shona",
			L_so:    "Somali",
			L_sq:    "Albanian",
			L_sr:    "Serbian",
			L_st:    "Southern Sotho",
			L_su:    "Sundanese",
			L_sv:    "Swedish",
			L_sw:    "Swahili",
			L_ta:    "Tamil",
			L_te:    "Telugu",
			L_tg:    "Tajik",
			L_th:    "Thai",
			L_tl:    "Tagalog",
			L_tr:    "Turkish",
			L_uk:    "Ukrainian",
			L_ur:    "Urdu",
			L_uz:    "Uzbek",
			L_vi:    "Vietnamese",
			L_xh:    "Xhosa",
			L_yi:    "Yiddish",
			L_yo:    "Yoruba",
			L_zh:    "Chinese",
			L_zh_CN: "Chinese_CN", // Chinese
			L_zh_TW: "Chinese_TW", // Chinese
			L_zu:    "zuZulu",
		}
	}
}

func initLangReseve() {
	if langList == nil {
		return
	}

	if langListR == nil {
		langListR = make(map[string]string)
	}

	for k, v := range langList {
		if langListR[v] != k {
			langListR[v] = k
		}
	}
}

func IsLang(value string) bool {
	l := len(value)
	if l <= wotoValues.BaseOneIndex || l > len(L_zh_CN) {
		return false
	}
	initLang()
	if len(langList[value]) == wotoValues.BaseIndex {
		return false
	}

	return true
}
