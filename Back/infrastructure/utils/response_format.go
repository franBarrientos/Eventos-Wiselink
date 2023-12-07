package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ParseValidationErrors(err error) string {
	var messages []string

	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		fieldName := e.Field()
		tagName := e.Tag()

		message := fmt.Sprintf("Field '%s' has an invalid value. '%s'", fieldName, tagName)
		messages = append(messages, message)
	}

	return strings.Join(messages, "\n")
}
