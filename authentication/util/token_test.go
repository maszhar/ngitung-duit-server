package util_test

import (
	"testing"

	"github.com/djeniusinvfest/inventora/auth/util"
)

func TestValidToken(t *testing.T) {
	key := "secret-keys"
	userId := "1234567890adsfaf"

	pl, err := util.CreateAccessToken(key, userId)
	if err != nil {
		t.Fatalf("%v", err)
	}

	subject, err := util.ParseAccessToken(key, pl)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if subject != userId {
		t.Fatalf("ParseAccessToken(token) = %v, wants %v", subject, userId)
	}
}
