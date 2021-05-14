package wotoTranslate

type Lang struct {
	Data LangData `json:"data"`
}

type LangData struct {
	Detections []LangDetect `json:"detections"`
}

type LangDetect struct {
	TheLang    string  `json:"language"`
	Reliable   bool    `json:"isReliable"`
	Confidence float32 `json:"confidence"`
}
