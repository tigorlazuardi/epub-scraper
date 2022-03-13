package logger

import "errors"

// Utility function that receive message that will return the message itself and error containing it's message.
func NewError(message string) (string, error) {
	return message, errors.New(message)
}
