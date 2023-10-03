package model

type Tchat struct {
	Id     int    `json:"id"`
	UUID   string `json:"uuid"`
	Author string `json:"author"`
	Msg    string `json:"msg"`
}
