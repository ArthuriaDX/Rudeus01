package wotoTranslate

type Lang struct {
	Data *LangData `json:"data"`
}

type LangData struct {
	Detections []LangDetect `json:"detections"`
}

type LangDetect struct {
	TheLang    string  `json:"language"`
	Reliable   bool    `json:"isReliable"`
	Confidence float32 `json:"confidence"`
}

type GoogleTr struct {
}

type gnuTranslate struct {
	Result string `json:"result"`
	Err    string `json:"error"`
}

type WotoTr struct {
	UserText       string
	OriginalText   string
	TranslatedText string
	From           string
	To             string
	CorrectedValue string
	HasWrongNess   bool
	WrongFrom      bool
	Road           map[int]bool
}
