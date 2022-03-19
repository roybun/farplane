package sessions

import "errors"

type FPSession struct {
	sessionToken string
	Username     string
}

type FPSessions interface {
	CreateSession(username string) (string, error)
	GetSession(sessionToken string) (FPSession, error)
}

var inMemSessions = [5]FPSession{}
var inMemSessionsCurIndex = 0

type InMemFPSessions struct{}

func (inmemsessions InMemFPSessions) GetSession(sessionToken string) (FPSession, error) {
	for index, fbSession := range inMemSessions {
		if "Bearer "+fbSession.sessionToken == sessionToken {
			return inMemSessions[index], nil
		}
	}
	return FPSession{}, errors.New("Session for user does not exist")
}
func (inmemsessions InMemFPSessions) CreateSession(username string) (string, error) {
	existingSession, error := inmemsessions.GetSession(username)
	if error != nil {
		inMemSessions[inMemSessionsCurIndex].Username = username
		inMemSessions[inMemSessionsCurIndex].sessionToken = "LeToken"
		inMemSessionsCurIndex++
		return inMemSessions[inMemSessionsCurIndex-1].sessionToken, nil
	} else {
		return existingSession.sessionToken, nil
	}
}
