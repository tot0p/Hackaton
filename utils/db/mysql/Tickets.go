package mysql

import (
	"github.com/google/uuid"
	"hackaton/model"
)

// CreateTicket creates a new ticket in the database
func CreateTicket(UserUUID, Subject, Content, LinkedData string) (*model.Ticket, error) {

	tk := model.Ticket{
		UUID:       uuid.New().String(),
		UserUUID:   UserUUID,
		Subject:    Subject,
		Content:    Content,
		Status:     model.StatusOpen,
		LinkedData: LinkedData,
		TchatUUID:  "",
	}

	// insert the ticket into the database
	_, err := DB.Conn.Exec("INSERT INTO tickets (uuid,user_uuid, subject, content, status, linked_data, tchat_uuid) VALUES (?,?, ?, ?, ?, ?, ?)", tk.UUID, tk.UserUUID, tk.Subject, tk.Content, tk.Status, tk.LinkedData, tk.TchatUUID)
	if err != nil {
		return nil, err
	}
	return &tk, nil
}
