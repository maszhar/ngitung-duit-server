package validator

import (
	"fmt"
	"strings"
)

func validateNotBlank(entity string, value string) error {
	if len(value) == 0 {
		return fmt.Errorf("validator: %s cannot be blank", entity)
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
