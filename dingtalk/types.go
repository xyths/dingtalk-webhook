package dingtalk

type Text struct {
	Content string `json:"content"`
}
type At struct {
	Mobiles []string `json:"atMobiles,omitempty"`
	AtAll   bool     `json:"isAtAll,omitempty"`
}
type MessageText struct {
	Type string `json:"msgtype"`
	Text Text   `json:"text"`
	At   *At    `json:"at,omitempty"`
}

type Response struct {
	Code    int64
	Message string
}
