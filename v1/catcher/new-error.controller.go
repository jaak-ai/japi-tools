package catcher

import "github.com/gofiber/fiber/v2/utils"

func NewError(httpCode int, message string, code int32, eventId string) *Error {
	err := &Error{
		Status:  httpCode,
		Message: utils.StatusMessage(httpCode),
		Code:    code,
		EventID: eventId,
	}

	if message != "" {
		err.Message = message
	}

	return err
}
