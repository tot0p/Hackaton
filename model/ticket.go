package model

const (
	StatusOpen = iota
	StatusClosed
	StatusInProgress
)

type Ticket struct {
	UUID       string `json:"uuid"`
	UserUUID   string `json:"user_uuid"`
	Subject    string `json:"subject"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
	LinkedData string `json:"linked_data"`
	CreatedAt  string `json:"created_at"`
	TchatUUID  string `json:"tchat_uuid"`
}
