package main

type Response struct {
	ResponseData ResponseData
	Rest         Rest
}

type ResponseData struct {
	TranslatedText string
	Match          float32
}

type Rest struct {
	QuotaFinished   bool
	MtLangSupported string
	ResponseDetails string
	ResponseStatus  int16
	ResponderId     string
	Exception_code  string
	Matches         []RestArray
}

type RestArray struct {
	Id             int
	Segment        string
	Translation    string
	Source         string
	Target         string
	Quality        int
	Reference      string
	UsageCount     int
	Subject        bool
	CreatedBy      string
	LastUpdatedBy  string
	CreateDate     string
	LastUpdateDate string
	Match          float32
	Model          string
}

type NotFoundEmoji struct {
	Message string `json:"Message,omitempty"`
}

type Emoji struct {
	Simbol string `json:"Simbol,omitempty"`
	Name   string `json:"Name,omitempty"`
}
