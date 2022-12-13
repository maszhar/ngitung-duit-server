package validator_test

import (
	"testing"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
	"github.com/djeniusinvfest/inventora/auth/validator"
	"github.com/jaswdr/faker"
)

func generateRegisterRequest() *pb.RegisterRequest {
	faker := faker.New()
	return &pb.RegisterRequest{
		FirstName: faker.Person().FirstName(),
		LastName:  faker.Person().LastName(),
		Email:     faker.Internet().Email(),
		Password:  faker.Internet().Password(),
		AgreeTos:  true,
	}
}

func TestRegisterFirstNameEmpty(t *testing.T) {
	in := generateRegisterRequest()
	in.FirstName = ""
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(first_name blank) not returns error, wants error")
	}
}

func TestRegisterLastNameEmpty(t *testing.T) {
	in := generateRegisterRequest()
	in.LastName = ""
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(last_name blank) not returns error, wants error")
	}
}

func TestRegisterEmailEmpty(t *testing.T) {
	in := generateRegisterRequest()
	in.Email = ""
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(email blank) not returns error, wants error")
	}
}

func TestRegisterEmailInvalid(t *testing.T) {
	in := generateRegisterRequest()
	in.Email = "invalidmail.com"
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(email invalid) not returns error, wants error")
	}
}

func TestRegisterPasswordBlank(t *testing.T) {
	in := generateRegisterRequest()
	in.Password = ""
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(password blank) not returns error, wants error")
	}
}

func TestRegisterPasswordBelow6Chars(t *testing.T) {
	in := generateRegisterRequest()
	in.Password = "jawab"
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(password below 6 characters) not returns error, wants error")
	}
}

func TestRegisterAgreeTosFalse(t *testing.T) {
	in := generateRegisterRequest()
	in.AgreeTos = false
	err := validator.ValidateRegister(in)
	if err == nil {
		t.Fatalf("RegisterValidator(agree_tos false) not returns error, wants error")
	}
}
