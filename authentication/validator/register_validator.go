package validator

import (
	"errors"
	"net/mail"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
)

func ValidateRegister(in *pb.RegisterRequest) error {
	trimSpacesArr([]*string{
		&in.FirstName,
		&in.LastName,
		&in.Email,
	})

	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return errors.New("validator: email must be valid")
	}

	err = validateNotBlank("first_name", in.FirstName)
	if err != nil {
		return err
	}

	err = validateNotBlank("last_name", in.LastName)
	if err != nil {
		return err
	}

	err = validateNotBlank("password", in.Password)
	if err != nil {
		return err
	}

	if len(in.Password) < 6 {
		return errors.New("validator: password length must be at least 6 characters")
	}

	if !in.AgreeTos {
		return errors.New("validator: agree_tos must be accepted")
	}

	return nil
}
