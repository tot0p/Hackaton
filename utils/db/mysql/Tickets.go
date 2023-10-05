package mysql

import (
	"errors"
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
		TchatUUID:  uuid.New().String(),
	}

	// insert the ticket into the database
	_, err := DB.Conn.Exec("INSERT INTO tickets (uuid,user_uuid, subject, content, status, linked_data, tchat_uuid) VALUES (?,?, ?, ?, ?, ?, ?)", tk.UUID, tk.UserUUID, tk.Subject, tk.Content, tk.Status, tk.LinkedData, tk.TchatUUID)
	if err != nil {
		return nil, err
	}
	return &tk, nil
}

// GetAllTickets gets all tickets from the database
func GetAllTickets() ([]*model.Ticket, error) {
	rows, err := DB.Conn.Query("SELECT * FROM tickets")
	if err != nil {
		return nil, err
	}
	var tickets []*model.Ticket
	for rows.Next() {
		tk := model.Ticket{}
		err := rows.Scan(&tk.UUID, &tk.UserUUID, &tk.Subject, &tk.Content, &tk.Status, &tk.LinkedData, &tk.CreatedAt, &tk.TchatUUID)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &tk)
	}
	return tickets, nil
}

// GetTicketByUUID gets a ticket from the database by uuid
func GetTicketByUUID(uuid string) (*model.Ticket, error) {
	tk := model.Ticket{}
	val := DB.Conn.QueryRow("SELECT * FROM tickets WHERE uuid = ?", uuid)
	err := val.Scan(&tk.UUID, &tk.UserUUID, &tk.Subject, &tk.Content, &tk.Status, &tk.LinkedData, &tk.CreatedAt, &tk.TchatUUID)
	if err != nil {
		return nil, err
	}
	return &tk, nil
}

// GetTicketsByUserUUID gets all tickets from the database by user uuid
func GetTicketsByUserUUID(userUUID string) ([]*model.Ticket, error) {
	rows, err := DB.Conn.Query("SELECT * FROM tickets WHERE user_uuid = ?", userUUID)
	if err != nil {
		return nil, err
	}
	var tickets []*model.Ticket
	for rows.Next() {
		tk := model.Ticket{}
		err := rows.Scan(&tk.UUID, &tk.UserUUID, &tk.Subject, &tk.Content, &tk.Status, &tk.LinkedData, &tk.CreatedAt, &tk.TchatUUID)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &tk)
	}
	return tickets, nil
}

func UpdateStatusTicket(uuid string, status int) error {
	if status > model.StatusClosed || status < model.StatusOpen {
		return errors.New("invalid status")
	}
	_, err := DB.Conn.Exec("UPDATE tickets SET status = ? WHERE uuid = ?", status, uuid)
	if err != nil {
		return err
	}
	return nil
}

func TchatToTchatUtilsInfo(tchats []*model.Tchat) []*model.TchatUtilsInfo {
	var tchatUtilsInfo []*model.TchatUtilsInfo

	var UUIDToName = make(map[string]string)

	for _, tchat := range tchats {
		if _, ok := UUIDToName[tchat.Author]; !ok {
			user, err := GetUserByUUID(tchat.Author)
			if err != nil {
				continue
			}
			UUIDToName[tchat.Author] = user.Username
		}
		tchatUtilsInfo = append(tchatUtilsInfo, &model.TchatUtilsInfo{
			Author: UUIDToName[tchat.Author],
			Msg:    tchat.Msg,
		})
	}

	return tchatUtilsInfo
}
