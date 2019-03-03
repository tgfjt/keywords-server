package models

type Keyword struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	User User   `json:"user"`
	Tags []*Tag `json:"tags"`
}

type NewKeyword struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewTag struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Tag struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	User User   `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
