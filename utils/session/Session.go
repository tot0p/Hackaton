package session

import (
	"github.com/google/uuid"
	"hackaton/model"
)

func init() {
	SessionsManager.Sessions = make(map[string]*session)
}

var SessionsManager sessionsManager

type sessionsManager struct {
	Sessions map[string]*session
}

type session struct {
	UUID string
	User *model.User
}

func GetSessionByUUID(uuid string) *session {
	return SessionsManager.Sessions[uuid]
}

func (s *sessionsManager) AddSession(user *model.User) *session {
	u := uuid.New().String()
	var ses *session = &session{
		UUID: u,
		User: user,
	}
	s.Sessions[u] = ses
	return ses
}

func (s *sessionsManager) RemoveSession(uuid string) {
	delete(s.Sessions, uuid)
}

func (s *session) GetUser() *model.User {
	return s.User
}

func (s *session) GetUUID() string {
	return s.UUID
}
