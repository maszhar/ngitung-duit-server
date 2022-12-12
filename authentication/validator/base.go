package validator

import (
	"errors"
	"fmt"
	"strings"
)

func validateNotBlank(entity string, value string) error {
	if len(value) == 0 {
		return errors.New(fmt.Sprintf("validator: %s cannot be blank", entity))
	}
	return nil
}

func trimSpaces(value *string) {
	*value = strings.Trim(*value, " ")
}

func trimSpacesArr(values []*string) {
	for _, value := range values {
		trimSpaces(value)
	}
}
