package model

// Tchat struct
type Tchat struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Msg     string `json:"msg"`
	Channel string `json:"channel"`
}

// TchatUtilsInfo struct
type TchatUtilsInfo struct {
	Msg    string `json:"msg"`
	Author string `json:"author"`
}
