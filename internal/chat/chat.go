package chat

type Chat struct {
	Messages []string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

var MsgHistory []Message
