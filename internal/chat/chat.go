package chat

type Chat struct {
	Messages []string
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

var MsgHistory []Messages
