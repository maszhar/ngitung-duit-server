package validator

import (
	"errors"
	"net/mail"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
)

func ValidateLogin(p *pb.LoginRequest) error {
	trimSpaces(&p.Email)

	_, err := mail.ParseAddress(p.Email)
	if err != nil {
		return errors.New("validator: email must be valid")
	}

	err = validateNotBlank("password", p.Password)
	if err != nil {
		return err
	}

	if len(p.Password) < 6 {
		return errors.New("validator: password length must be at least 6 characters")
	}

	return nil
}
