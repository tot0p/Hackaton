package model

type Tchat struct {
	UUID    string `json:"uuid"`
	Author  string `json:"author"`
	Msg     string `json:"msg"`
	Channel string `json:"channel"`
	ReplyOf string `json:"reply_of"`
}

type CompleteTchat struct {
	Author string `json:"author"`
	Msg    string `json:"msg"`
}
