package mysql

import (
	"github.com/google/uuid"
	"hackaton/model"
)

func CreateTchat(UserUUID, Msg, Channel string) (*model.Tchat, error) {

	ReplyOf := ""

	// Get uuid of the last message of the channel
	val := DB.Conn.QueryRow("SELECT uuid FROM tchat WHERE channel = ? ORDER BY uuid DESC LIMIT 1", Channel)
	err := val.Scan(&ReplyOf)
	if err != nil {
		ReplyOf = ""
	}

	tk := model.Tchat{
		UUID:    uuid.New().String(),
		Author:  UserUUID,
		Msg:     Msg,
		ReplyOf: ReplyOf,
		Channel: Channel,
	}

	// insert the ticket into the database
	_, err = DB.Conn.Exec("INSERT INTO tchat (uuid,author, msg, reply_of , channel) VALUES (?,?, ?, ?, ?)", tk.UUID, tk.Author, tk.Msg, tk.ReplyOf, tk.Channel)
	if err != nil {
		return nil, err
	}

	return &tk, nil

}

func GetTchatByChannel(Channel string) ([]*model.Tchat, error) {
	rows, err := DB.Conn.Query("SELECT * FROM tchat WHERE channel = ?", Channel)
	if err != nil {
		return nil, err
	}
	var tchats []*model.Tchat
	for rows.Next() {
		tk := model.Tchat{}
		err := rows.Scan(&tk.UUID, &tk.Author, &tk.Msg, &tk.Channel, &tk.ReplyOf)
		if err != nil {
			return nil, err
		}
		tchats = append(tchats, &tk)
	}
	return tchats, nil
}
