package validator_test

import (
	"testing"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/validator"
	"github.com/jaswdr/faker"
)

func generateLoginRequest() *pb.LoginRequest {
	faker := faker.New()
	return &pb.LoginRequest{
		Email:    faker.Internet().Email(),
		Password: faker.Internet().Password(),
	}
}

func TestLoginEmailBlank(t *testing.T) {
	p := generateLoginRequest()
	p.Email = ""
	err := validator.ValidateLogin(p)
	if err == nil {
		t.Fatalf("LoginValidator(email blank) = %v, wants error", err)
	}
}

func TestLoginEmailInvalid(t *testing.T) {
	p := generateLoginRequest()
	p.Email = "invalid-mail"
	err := validator.ValidateLogin(p)
	if err == nil {
		t.Fatalf("LoginValidator(email invalid) = %v, wants error", err)
	}
}

func TestLoginPasswordBlank(t *testing.T) {
	p := generateLoginRequest()
	p.Password = ""
	err := validator.ValidateLogin(p)
	if err == nil {
		t.Fatalf("LoginValidator(password blank) = %v, wants error", err)
	}
}

func TestLoginPasswordShort(t *testing.T) {
	p := generateLoginRequest()
	p.Password = "koala"
	err := validator.ValidateLogin(p)
	if err == nil {
		t.Fatalf("LoginValidator(password blank) = %v, wants error", err)
	}
}
