package sessions

import (
	"testing"
)

func TestInMemCreateSession(t *testing.T) {
	inmemsessions := InMemFPSessions{}

	leToken, err := inmemsessions.CreateSession("User1")
	if err != nil {
		t.Fatalf("Error returned attempting to create session")
	}

	leSession, err := inmemsessions.GetSession(leToken)
	if err != nil {
		t.Fatalf("Error attempting to get session")
	}

	if leSession.sessionToken != leToken {
		t.Fatalf("Tokens between create and get do not match")
	}
}
