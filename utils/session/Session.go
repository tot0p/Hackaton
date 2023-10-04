package session

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hackaton/model"
)

func init() {
	SessionsManager.Sessions = make(map[string]*Session)
}

var SessionsManager sessionsManager

type sessionsManager struct {
	Sessions map[string]*Session
}

type Session struct {
	UUID string
	User *model.User
}

func (s *sessionsManager) IsLogged(ctx *gin.Context) bool {
	cookie, err := ctx.Cookie("session")
	if err != nil {
		return false
	}
	_, ok := s.Sessions[cookie]
	return ok
}

func (s *sessionsManager) GetUser(ctx *gin.Context) *model.User {
	cookie, err := ctx.Cookie("session")
	if err != nil {
		return nil
	}
	return s.Sessions[cookie].GetUser()
}

func GetSessionByUUID(uuid string) *Session {
	return SessionsManager.Sessions[uuid]
}

func (s *sessionsManager) CreateSession(ctx *gin.Context, user *model.User) {
	ses := s.AddSession(user)
	domain := ctx.Request.Host
	ctx.SetCookie("session", ses.GetUUID(), 3600, "/", domain, false, false)
}

func (s *sessionsManager) AddSession(user *model.User) *Session {
	u := uuid.New().String()
	var ses = &Session{
		UUID: u,
		User: user,
	}
	s.Sessions[u] = ses
	return ses
}

func (s *sessionsManager) RemoveSession(uuid string) {
	delete(s.Sessions, uuid)
}

func (s *Session) GetUser() *model.User {
	return s.User
}

func (s *Session) GetUUID() string {
	return s.UUID
}
