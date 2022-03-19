package authentication

import (
	"testing"
)

func TestHardcodedValidateUserOk(t *testing.T) {
	hca := HardcodedAuthenticator{}
	err := hca.ValidateUser("User1", "PW1")
	if err != nil {
		t.Fatalf("Could not login User1 with PW1")
	}
}

func TestHardcodedValidateUserErr(t *testing.T) {
	hca := HardcodedAuthenticator{}
	err := hca.ValidateUser("User1", "badpw")
	if err == nil {
		t.Fatalf("")
	}
}
